package interactor

import "clean_architecture/app/domains/entity"

type UserInteractor struct {
	repository *UserRepository
	presenter  *UserPresenter
}

func NewUserInteractor(user *UserRepository, presenter *UserPresenter) *UserInteractor {
	return &UserInteractor{
		repository: user,
		presenter:  presenter,
	}
}

func (i *UserInteractor) ListUsers() (data []byte, err error) {
	users, err := (*i.repository).GetAllUsers()
	if err != nil {
		return
	}
	inp := newDSUsersPresenter(users)
	data, err = (*i.presenter).ListUsersView(inp)
	return
}

func (i *UserInteractor) CreateUser(obj *entity.User) (data []byte, err error) {
	user, err := (*i.repository).CreateUser(obj)
	if err != nil {
		return
	}
	inp := newDSUserPresenter(user)
	data, err = (*i.presenter).CreateUserView(inp)
	return
}

func (i *UserInteractor) RetrieveUser(id int) (data []byte, err error) {
	user, err := (*i.repository).FindUserById(id)
	if err != nil {
		return
	}
	inp := newDSUserPresenter(user)
	data, err = (*i.presenter).RetrieveUserView(inp)
	return
}

func (i *UserInteractor) UpdateUser(id int, obj *entity.User) (data []byte, err error) {
	user, err := (*i.repository).UpdateUserById(id, obj)
	if err != nil {
		return
	}
	inp := newDSUserPresenter(user)
	data, err = (*i.presenter).UpdateUserView(inp)
	return
}

func (i *UserInteractor) DeleteUser(id int) (data []byte, err error) {
	user, err := (*i.repository).DeleteUserById(id)
	if err != nil {
		return
	}
	inp := newDSUserPresenter(user)
	data, err = (*i.presenter).DeleteUserView(inp)
	return
}
