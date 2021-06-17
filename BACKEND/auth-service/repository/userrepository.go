package repository

import (
	"fmt"
	"xml/auth-service/data"

	"gorm.io/gorm"
)
type UserRepository struct {
	Database *gorm.DB
}

func (repo *UserRepository) CreateUser(user *data.User) error {
	result := repo.Database.Create(user)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *UserRepository) UserExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.User{}).Count(&count)
	return count != 0
}

func (repo *UserRepository) FindUserByID(id uint) *data.User {
	var user data.User
	repo.Database.Where("id = ?", id).First(&user)
	return &user
}

func(repo *UserRepository) UpdateUser(user *data.User) error {
	err := repo.Database.Save(&user).Error

	return err
}

func(repo *UserRepository) UserExistsByMail(email string) bool {
	var count int64
	repo.Database.Where("email = ?", email).Find(&data.User{}).Count(&count)
	return count != 0
}

func(repo *UserRepository) UserExistsByUsername(username string) bool {
	var count int64
	repo.Database.Where("username = ?", username).Find(&data.User{}).Count(&count)
	return count != 0
}