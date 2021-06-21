package service

import (
	"fmt"
	"xml/auth-service/data"
	"xml/auth-service/dto"
	"xml/auth-service/repository"
	"xml/auth-service/security"
)

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

func (service *UserService) FindUserByUsername(username string) *data.User {
	var user data.User
	service.Repo.Database.Where("username = ?", username).First(&user)
	return &user
}

func (service *UserService) EditUserData(userEditDTO dto.UserEditDTO, oldUsername string) error {
	user := service.FindUserByUsername(oldUsername)

	if user == nil {
		return nil
	}

	if oldUsername != userEditDTO.Username {
		if service.Repo.UserExistsByUsername(userEditDTO.Username) {
			return fmt.Errorf("Username is taken")
		}
	}

	if user.Email != userEditDTO.Email {
		if service.Repo.UserExistsByMail(userEditDTO.Email) {
			return fmt.Errorf("Email is taken")
		}
	}

	user.Username = userEditDTO.Username
	user.Email = userEditDTO.Email
	user.Name = userEditDTO.Name
	user.LastName = userEditDTO.LastName

	err := service.Repo.UpdateUser(user)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func(service *UserService) ChangePassword(username string, newpassword string) error {

	user := service.FindUserByUsername(username)

	newpw, err := security.HashPassword(newpassword)

	if err != nil {
		fmt.Println(err)
		return err
	}

	user.Password = newpw
	err = service.Repo.UpdateUser(user)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func(service *UserService) GetCurrentUser(username string) *data.User {
	user := service.FindUserByUsername(username)

	return user
}

func (service *UserService) GetIDByUsername(username string) uint {
	userId := service.Repo.FindIDByUsername(username)

	return userId
}

