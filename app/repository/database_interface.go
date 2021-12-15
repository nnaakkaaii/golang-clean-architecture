package repository

import "clean_architecture/app/interactor"

type UserDatabaseInterface interface {
	ListUsers() (users []*interactor.UserOut, err error)
	CreateUser(obj *interactor.UserIn) (id int64, err error)
	RetrieveUser(pk int) (user *interactor.UserOut, err error)
	UpdateUser(pk int, obj *interactor.UserIn) (err error)
	DeleteUser(pk int) (err error)
}
