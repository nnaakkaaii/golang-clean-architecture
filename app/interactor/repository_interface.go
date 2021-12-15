package interactor

type UserRepositoryInterface interface {
	GetAllUsers() (users []*UserOut, err error)
	CreateUser(obj *UserIn) (user *UserOut, err error)
	FindUserById(pk int) (user *UserOut, err error)
	UpdateUserById(pk int, obj *UserIn) (user *UserOut, err error)
	DeleteUserById(pk int) (user *UserOut, err error)
}
