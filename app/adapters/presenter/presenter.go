package presenter

import (
	"clean_architecture/app/usecases/interactor"
	"fmt"
)

type UserPresenter struct {
	view *UserView
}

var _ interactor.UserPresenter = (*UserPresenter)(nil)

func NewUserPresenter(view *UserView) *UserPresenter {
	return &UserPresenter{view: view}
}

func (presenter *UserPresenter) ListUsersView(users []*interactor.DSUserPresenter) (data []byte, err error) {
	res := newResponseUsersView(users)
	data, err = (*presenter.view).ListUsersView(res)
	if err != nil {
		return
	}
	return
}

func (presenter *UserPresenter) CreateUserView(user *interactor.DSUserPresenter) (data []byte, err error) {
	res := newResponseUserView(user)
	data, err = (*presenter.view).DeleteUserView(res)
	if err != nil {
		return
	}
	return
}

func (presenter *UserPresenter) RetrieveUserView(user *interactor.DSUserPresenter) (data []byte, err error) {
	res := newResponseUserView(user)
	data, err = (*presenter.view).RetrieveUserView(res)
	if err != nil {
		return
	}
	return
}

func (presenter *UserPresenter) UpdateUserView(user *interactor.DSUserPresenter) (data []byte, err error) {
	res := newResponseUserView(user)
	data, err = (*presenter.view).UpdateUserView(res)
	if err != nil {
		return
	}
	return
}

func (presenter *UserPresenter) DeleteUserView(user *interactor.DSUserPresenter) (data []byte, err error) {
	res := newResponseUserView(user)
	data, err = (*presenter.view).DeleteUserView(res)
	if err != nil {
		return
	}
	return
}

func formatFullName(firstName, lastName string) (fullName string) {
	fullName = fmt.Sprintf("%s %s", firstName, lastName)
	return
}

func formatBirthDate(birthYear, birthMonth, birthDay int) (birthDate string) {
	birthDate = fmt.Sprintf("%04d/%02d/%02d", birthYear, birthMonth, birthDay)
	return
}
