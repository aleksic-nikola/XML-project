package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/monolit-service/data"
)

type UserRepository struct {
	Database *gorm.DB
}

func (repo *UserRepository) CreateUser(user *data.User) error {
	result := repo.Database.Create(user)
	fmt.Println(result.RowsAffected)
	return result.Error
}