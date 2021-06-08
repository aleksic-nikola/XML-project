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