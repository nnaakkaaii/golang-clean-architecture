package view

import (
	"clean_architecture/app/adapters/presenter"
	"encoding/json"
)

type UserView struct {
}

var _ presenter.UserView = (*UserView)(nil)

func NewView() *UserView {
	return &UserView{}
}

func (v *UserView) ListUsersView(response *presenter.ResponseUsersView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func (v *UserView) CreateUserView(response *presenter.ResponseUserView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func (v *UserView) RetrieveUserView(response *presenter.ResponseUserView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func (v *UserView) UpdateUserView(response *presenter.ResponseUserView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func (v *UserView) DeleteUserView(response *presenter.ResponseUserView) (data []byte, err error) {
	data, err = jsonify(response)
	return
}

func jsonify(body interface{}) (data []byte, err error) {
	data, err = json.Marshal(body)
	return
}
