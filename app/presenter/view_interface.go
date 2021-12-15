package presenter

type UserViewInterface interface {
	ListUsersView(res *UsersView) (data []byte, err error)
	CreateUserView(res *UserView) (data []byte, err error)
	RetrieveUserView(res *UserView) (data []byte, err error)
	UpdateUserView(res *UserView) (data []byte, err error)
	DeleteUserView(res *UserView) (data []byte, err error)
}
