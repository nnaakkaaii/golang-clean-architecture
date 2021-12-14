package interactor

import "clean_architecture/app/domains/entity"

type UserRepository interface {
	GetAllUsers() (users []*entity.User, err error)
	CreateUser(obj *entity.User) (user *entity.User, err error)
	FindUserById(pk int) (user *entity.User, err error)
	UpdateUserById(pk int, obj *entity.User) (user *entity.User, err error)
	DeleteUserById(pk int) (user *entity.User, err error)
}
