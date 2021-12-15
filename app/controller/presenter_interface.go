package controller

import "clean_architecture/app/interactor"

type UserPresenterInterface interface {
	ListUsersView(users []*interactor.UserResponse) (data []byte, err error)
	CreateUserView(user *interactor.UserResponse) (data []byte, err error)
	RetrieveUserView(user *interactor.UserResponse) (data []byte, err error)
	UpdateUserView(user *interactor.UserResponse) (data []byte, err error)
	DeleteUserView(user *interactor.UserResponse) (data []byte, err error)
}
