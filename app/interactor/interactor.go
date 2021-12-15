package interactor

type UserInteractor struct {
	repository *UserRepositoryInterface
}

func NewUserInteractor(user *UserRepositoryInterface) *UserInteractor {
	return &UserInteractor{
		repository: user,
	}
}

var _ UserInteractorInterface = (*UserInteractor)(nil)

func (i *UserInteractor) ListUsers() (data []*UserResponse, err error) {
	out, err := (*i.repository).GetAllUsers()
	if err != nil {
		return
	}
	for _, o := range out {
		d := UserResponse(*o)
		data = append(data, &d)
	}
	return
}

func (i *UserInteractor) CreateUser(obj *UserRequest) (data *UserResponse, err error) {
	in := UserIn(*obj)
	out, err := (*i.repository).CreateUser(&in)
	if err != nil {
		return
	}
	res := UserResponse(*out)
	data = &res
	return
}

func (i *UserInteractor) RetrieveUser(id int) (data *UserResponse, err error) {
	out, err := (*i.repository).FindUserById(id)
	if err != nil {
		return
	}
	res := UserResponse(*out)
	data = &res
	return
}

func (i *UserInteractor) UpdateUser(id int, obj *UserRequest) (data *UserResponse, err error) {
	in := UserIn(*obj)
	out, err := (*i.repository).UpdateUserById(id, &in)
	if err != nil {
		return
	}
	res := UserResponse(*out)
	data = &res
	return
}

func (i *UserInteractor) DeleteUser(id int) (data *UserResponse, err error) {
	out, err := (*i.repository).DeleteUserById(id)
	if err != nil {
		return
	}
	res := UserResponse(*out)
	data = &res
	return
}
