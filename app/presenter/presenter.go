package presenter

import (
	"clean_architecture/app/controller"
	"clean_architecture/app/interactor"
)

type UserPresenter struct {
	view *UserViewInterface
}

var _ controller.UserPresenterInterface = (*UserPresenter)(nil)

func NewUserPresenter(view *UserViewInterface) *UserPresenter {
	return &UserPresenter{view: view}
}

func (presenter *UserPresenter) ListUsersView(users []*interactor.UserResponse) (data []byte, err error) {
	res := newUsersView(users)
	data, err = (*presenter.view).ListUsersView(res)
	return
}

func (presenter *UserPresenter) CreateUserView(user *interactor.UserResponse) (data []byte, err error) {
	res := newUserView(user)
	data, err = (*presenter.view).CreateUserView(res)
	return
}

func (presenter *UserPresenter) RetrieveUserView(user *interactor.UserResponse) (data []byte, err error) {
	res := newUserView(user)
	data, err = (*presenter.view).RetrieveUserView(res)
	return
}

func (presenter *UserPresenter) UpdateUserView(user *interactor.UserResponse) (data []byte, err error) {
	res := newUserView(user)
	data, err = (*presenter.view).UpdateUserView(res)
	return
}

func (presenter *UserPresenter) DeleteUserView(user *interactor.UserResponse) (data []byte, err error) {
	res := newUserView(user)
	data, err = (*presenter.view).DeleteUserView(res)
	return
}
