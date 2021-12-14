package presenter

import "clean_architecture/app/usecases/interactor"

type ResponseUsersView struct {
	Users []*ResponseUser `json:"users"`
}

func newResponseUsersView(users []*interactor.DSUserPresenter) *ResponseUsersView {
	res := make([]*ResponseUser, 0, len(users))
	for _, user := range users {
		r := ResponseUser(*newResponseUserView(user))
		res = append(res, &r)
	}
	return &ResponseUsersView{Users: res}
}

type ResponseUserView ResponseUser

type ResponseUser struct {
	Id        string `json:"id"`
	FullName  string `json:"full_name"`
	BirthDate string `json:"birth_date"`
}

func newResponseUserView(user *interactor.DSUserPresenter) *ResponseUserView {
	return &ResponseUserView{
		Id:        user.Id,
		FullName:  formatFullName(user.FirstName, user.LastName),
		BirthDate: formatBirthDate(user.BirthYear, user.BirthMonth, user.BirthDay),
	}
}
