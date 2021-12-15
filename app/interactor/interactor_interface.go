package interactor

type UserInteractorInterface interface {
	ListUsers() (data []*UserResponse, err error)
	CreateUser(obj *UserRequest) (data *UserResponse, err error)
	RetrieveUser(id int) (data *UserResponse, err error)
	UpdateUser(id int, obj *UserRequest) (data *UserResponse, err error)
	DeleteUser(id int) (data *UserResponse, err error)
}
