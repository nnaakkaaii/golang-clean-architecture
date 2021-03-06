package view

import (
	"clean_architecture/app/presenter"
	"encoding/json"
)

type UserView struct {
}

var _ presenter.UserViewInterface = (*UserView)(nil)

func NewUserView() *UserView {
	return &UserView{}
}

func (v *UserView) ListUsersView(response *presenter.UsersView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func (v *UserView) CreateUserView(response *presenter.UserView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func (v *UserView) RetrieveUserView(response *presenter.UserView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func (v *UserView) UpdateUserView(response *presenter.UserView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func (v *UserView) DeleteUserView(response *presenter.UserView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func jsonify(body interface{}) (data []byte, err error) {
	data, err = json.Marshal(body)
	return
}
