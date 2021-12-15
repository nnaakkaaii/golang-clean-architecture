package database

import (
	"clean_architecture/app/interactor"
	"clean_architecture/app/repository"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type UserDatabase struct {
	handler *sql.DB
}

var _ repository.UserDatabaseInterface = (*UserDatabase)(nil)

func NewUserDatabase(dbUser, dbPass, dbHost, dbPort, dbName string) *UserDatabase {
	ds := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	connector, err := sql.Open("postgres", ds)
	if err != nil {
		panic(err.Error())
	}
	return &UserDatabase{handler: connector}
}

func (database *UserDatabase) ListUsers() (users []*interactor.UserOut, err error) {
	query := `SELECT * FROM USERS`
	rows, err := database.handler.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var id, firstName, lastName string
	var birthDate time.Time
	for rows.Next() {
		if err := rows.Scan(&id, &firstName, &lastName, &birthDate); err != nil {
			return nil, err
		}
		users = append(users, &interactor.UserOut{
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

func (database *UserDatabase) CreateUser(obj *interactor.UserIn) (id int64, err error) {
	query := fmt.Sprintf(`INSERT INTO USERS(first_name, last_name, birth_date) VALUES('%s', '%s', '%s') RETURNING id`, obj.FirstName, obj.LastName, obj.FormatDate())
	log.Print(query)
	err = database.handler.QueryRow(query).Scan(&id)
	return
}

func (database *UserDatabase) RetrieveUser(pk int) (user *interactor.UserOut, err error) {
	var id, firstName, lastName string
	var birthDate time.Time
	query := fmt.Sprintf(`SELECT * FROM USERS WHERE id = %d LIMIT 1`, pk)
	if err := database.handler.QueryRow(query).Scan(&id, &firstName, &lastName, &birthDate); err != nil {
		return nil, err
	}
	user = &interactor.UserOut{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		BirthYear:  birthDate.Year(),
		BirthMonth: int(birthDate.Month()),
		BirthDay:   birthDate.Day(),
	}
	return
}

func (database *UserDatabase) UpdateUser(pk int, obj *interactor.UserIn) (err error) {
	query := fmt.Sprintf(`UPDATE USERS SET first_name='%s', last_name='%s', birth_date='%s' WHERE id=%d RETURNING id`, obj.FirstName, obj.LastName, obj.FormatDate(), pk)
	log.Print(query)
	err = database.handler.QueryRow(query).Err()
	return
}

func (database *UserDatabase) DeleteUser(pk int) (err error) {
	query := fmt.Sprintf(`DELETE FROM USERS WHERE id=%d`, pk)
	err = database.handler.QueryRow(query).Err()
	return
}
