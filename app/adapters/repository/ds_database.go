package repository

import "clean_architecture/app/domains/entity"

// DSUserDatabaseOutput はinterface層のdatabaseからadapter層のrepositoryに入力されるデータで、
// repositoryにとって都合の良いデータ構造であるべき (=databaseの事情に依存しない)
type DSUserDatabaseOutput struct {
	Id         string
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}

func (o *DSUserDatabaseOutput) convert() *entity.User {
	date := *entity.NewDate(o.BirthYear, o.BirthMonth, o.BirthDay)
	return entity.NewUser(o.Id, o.FirstName, o.LastName, date)
}

type DSUserDatabaseInput struct {
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}

func newDSUserDatabaseInput(obj *entity.User) *DSUserDatabaseInput {
	date := obj.GetBirthDate()
	return &DSUserDatabaseInput{
		FirstName:  obj.GetFirstName(),
		LastName:   obj.GetLastName(),
		BirthYear:  date.GetYear(),
		BirthMonth: date.GetMonth(),
		BirthDay:   date.GetDay(),
	}
}
