package database

import (
	"clean_architecture/app/adapters/repository"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type UserDatabase struct {
	handler *sql.DB
}

var _ repository.UserDatabase = (*UserDatabase)(nil)

func NewUserDatabase(dbUser, dbPass, dbHost, dbPort, dbName string) *UserDatabase {
	connector, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName))
	if err != nil {
		panic(err.Error())
	}
	return &UserDatabase{handler: connector}
}

func (database *UserDatabase) ListUsers() (users []*repository.DSUserDatabaseOutput, err error) {
	query := "SELECT * FROM USERS"
	rows, err := database.handler.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, firstName, lastName string
		var birthDate time.Time
		if err := rows.Scan(&id, &firstName, &lastName, &birthDate); err != nil {
			return nil, err
		}
		users = append(users, &repository.DSUserDatabaseOutput{
			Id:         id,
			FirstName:  firstName,
			LastName:   lastName,
			BirthYear:  birthDate.Year(),
			BirthMonth: int(birthDate.Month()),
			BirthDay:   birthDate.Day(),
		})
	}
	return
}

func (database *UserDatabase) CreateUser(obj *repository.DSUserDatabaseInput) (id int64, err error) {
	query := fmt.Sprintf(`INSERT INTO USERS(first_name, last_name, birth_date) VALUES('%s', '%s', '%04d/%02d/%02d') RETURNING id`, obj.FirstName, obj.LastName, obj.BirthYear, obj.BirthMonth, obj.BirthDay)
	err = database.handler.QueryRow(query).Scan(&id)
	return
}

func (database *UserDatabase) RetrieveUser(pk int) (user *repository.DSUserDatabaseOutput, err error) {
	var id, firstName, lastName string
	var birthDate time.Time
	query := fmt.Sprintf(`SELECT * FROM USERS WHERE id = %d LIMIT 1`, pk)
	if err := database.handler.QueryRow(query).Scan(&id, &firstName, &lastName, &birthDate); err != nil {
		return nil, err
	}
	user = &repository.DSUserDatabaseOutput{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		BirthYear:  birthDate.Year(),
		BirthMonth: int(birthDate.Month()),
		BirthDay:   birthDate.Day(),
	}
	return
}

func (database *UserDatabase) UpdateUser(pk int, obj *repository.DSUserDatabaseInput) (err error) {
	query := fmt.Sprintf(`UPDATE USERS SET first_name='%s', last_name='%s', birth_date='%04d/%02d/%02d' WHERE id=%d RETURNING id`, obj.FirstName, obj.LastName, obj.BirthYear, obj.BirthMonth, obj.BirthDay, pk)
	err = database.handler.QueryRow(query).Err()
	return
}

func (database *UserDatabase) DeleteUser(pk int) (err error) {
	query := fmt.Sprintf(`DELETE FROM USERS WHERE id=%d`, pk)
	err = database.handler.QueryRow(query).Err()
	return
}
