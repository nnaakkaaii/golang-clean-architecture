package presenter

import (
	"clean_architecture/app/interactor"
	"fmt"
)

type userView struct {
	Id        string `json:"id"`
	FullName  string `json:"full_name"`
	BirthDate string `json:"birth_date"`
}

type UserView userView

func newUserView(user *interactor.UserResponse) *UserView {
	return &UserView{
		Id:        user.Id,
		FullName:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		BirthDate: fmt.Sprintf("%04d/%02d/%02d", user.BirthYear, user.BirthMonth, user.BirthDay),
	}
}

type UsersView struct {
	Users []userView
}

func newUsersView(users []*interactor.UserResponse) *UsersView {
	res := make([]userView, 0, len(users))
	for _, user := range users {
		res = append(res, userView(*newUserView(user)))
	}
	return &UsersView{Users: res}
}
