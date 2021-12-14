package interactor

import "clean_architecture/app/domains/entity"

// DSUserPresenter はusecase層のinteractorからadapter層のpresenterに渡されるデータで、
// interactorに取って都合の良いデータ構造であり (=presenterの事情に依存しない)
// presenterにとっても描画に十分な情報であるべき
type DSUserPresenter struct {
	Id         string
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}

// newDSUserPresenter はentityが持つデータ構造である entity.User からpresenterに渡す際のデータ構造である DSUserPresenter に変換する
func newDSUserPresenter(user *entity.User) *DSUserPresenter {
	birthDate := user.GetBirthDate()
	return &DSUserPresenter{
		Id:         user.GetId(),
		FirstName:  user.GetFirstName(),
		LastName:   user.GetLastName(),
		BirthYear:  birthDate.GetYear(),
		BirthMonth: birthDate.GetMonth(),
		BirthDay:   birthDate.GetDay(),
	}
}

// newDSUsersPresenter は []entity.User から []DSUserPresenter を生成するための薄いラッパー
func newDSUsersPresenter(users []*entity.User) []*DSUserPresenter {
	ret := make([]*DSUserPresenter, 0, len(users))
	for _, user := range users {
		r := newDSUserPresenter(user)
		ret = append(ret, r)
	}
	return ret
}
