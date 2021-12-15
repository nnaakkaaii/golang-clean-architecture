package repository

import (
	"clean_architecture/app/interactor"
)

type UserRepository struct {
	database *UserDatabaseInterface
}

var _ interactor.UserRepositoryInterface = (*UserRepository)(nil)

func NewUserRepository(database *UserDatabaseInterface) *UserRepository {
	return &UserRepository{database: database}
}

func (repository *UserRepository) GetAllUsers() (users []*interactor.UserOut, err error) {
	users, err = (*repository.database).ListUsers()
	return
}

func (repository *UserRepository) CreateUser(obj *interactor.UserIn) (user *interactor.UserOut, err error) {
	id, err := (*repository.database).CreateUser(obj)
	if err != nil {
		return
	}
	user, err = (*repository.database).RetrieveUser(int(id))
	return
}

func (repository *UserRepository) FindUserById(pk int) (user *interactor.UserOut, err error) {
	user, err = (*repository.database).RetrieveUser(pk)
	return
}

func (repository *UserRepository) UpdateUserById(pk int, obj *interactor.UserIn) (user *interactor.UserOut, err error) {
	if err = (*repository.database).UpdateUser(pk, obj); err != nil {
		return
	}
	user, err = (*repository.database).RetrieveUser(pk)
	return
}

func (repository *UserRepository) DeleteUserById(pk int) (user *interactor.UserOut, err error) {
	user, err = (*repository.database).RetrieveUser(pk)
	if err != nil {
		return nil, err
	}
	err = (*repository.database).DeleteUser(pk)
	return
}
