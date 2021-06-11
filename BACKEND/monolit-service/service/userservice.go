package service

import (
	"xml/monolit-service/data"
	"xml/monolit-service/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (service *UserService) CreateUser(user *data.User) error {
	error := service.Repo.CreateUser(user)
	return error
}