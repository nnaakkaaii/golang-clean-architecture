package interactor

type UserPresenter interface {
	ListUsersView(users []*DSUserPresenter) (data []byte, err error)
	CreateUserView(user *DSUserPresenter) (data []byte, err error)
	RetrieveUserView(user *DSUserPresenter) (data []byte, err error)
	UpdateUserView(user *DSUserPresenter) (data []byte, err error)
	DeleteUserView(user *DSUserPresenter) (data []byte, err error)
}
