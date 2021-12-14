package presenter

type UserView interface {
	ListUsersView(response *ResponseUsersView) (data []byte, err error)
	CreateUserView(response *ResponseUserView) (data []byte, err error)
	RetrieveUserView(response *ResponseUserView) (data []byte, err error)
	UpdateUserView(response *ResponseUserView) (data []byte, err error)
	DeleteUserView(response *ResponseUserView) (data []byte, err error)
}
