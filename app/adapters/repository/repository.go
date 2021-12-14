package repository

import (
	"clean_architecture/app/domains/entity"
	"clean_architecture/app/usecases/interactor"
)

type UserRepository struct {
	database *UserDatabase
}

var _ interactor.UserRepository = (*UserRepository)(nil)

func NewUserRepository(database *UserDatabase) *UserRepository {
	return &UserRepository{database: database}
}

func (repository *UserRepository) GetAllUsers() (users []*entity.User, err error) {
	data, err := (*repository.database).ListUsers()
	if err != nil {
		return
	}
	for _, d := range data {
		users = append(users, d.convert())
	}
	return
}

func (repository *UserRepository) CreateUser(obj *entity.User) (user *entity.User, err error) {
	inp := newDSUserDatabaseInput(obj)
	id, err := (*repository.database).CreateUser(inp)
	if err != nil {
		return
	}
	d, err := (*repository.database).RetrieveUser(int(id))
	if err != nil {
		return
	}
	user = d.convert()
	return
}

func (repository *UserRepository) FindUserById(pk int) (user *entity.User, err error) {
	d, err := (*repository.database).RetrieveUser(pk)
	if err != nil {
		return
	}
	user = d.convert()
	return
}

func (repository *UserRepository) UpdateUserById(pk int, obj *entity.User) (user *entity.User, err error) {
	inp := newDSUserDatabaseInput(obj)
	if err = (*repository.database).UpdateUser(pk, inp); err != nil {
		return
	}
	d, err := (*repository.database).RetrieveUser(pk)
	if err != nil {
		return
	}
	user = d.convert()
	return
}

func (repository *UserRepository) DeleteUserById(pk int) (user *entity.User, err error) {
	d, err := (*repository.database).RetrieveUser(pk)
	if err != nil {
		return
	}
	if err = (*repository.database).DeleteUser(pk); err != nil {
		return
	}
	user = d.convert()
	return
}
