package service

import (
	"xml/auth-service/data"
	"xml/auth-service/repository")


type UserService struct {
	Repo *repository.UserRepository
}

func (service *UserService) CreateUser(user *data.User) error {
	error := service.Repo.CreateUser(user)
	return error
}

func (service *UserService) UserExists(id uint) (bool, error) {
	
	exists := service.Repo.UserExists(id)
	return exists, nil
}