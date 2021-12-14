package repository

// UserDatabase はinterfaces層のdatabaseに対するインターフェイスであり、
// databaseに何を採用するかを意識しなくて良い
type UserDatabase interface {
	ListUsers() (users []*DSUserDatabaseOutput, err error)
	CreateUser(obj *DSUserDatabaseInput) (id int64, err error)
	RetrieveUser(pk int) (user *DSUserDatabaseOutput, err error)
	UpdateUser(pk int, obj *DSUserDatabaseInput) (err error)
	DeleteUser(pk int) (err error)
}
