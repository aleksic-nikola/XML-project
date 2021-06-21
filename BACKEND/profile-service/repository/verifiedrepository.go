package repository

import (
	"fmt"
	"xml/profile-service/data"

	"gorm.io/gorm"
)
type VerifiedRepository struct {
	Database *gorm.DB
}

func (repo *VerifiedRepository) CreateVerified(user *data.Verified) error {
	result := repo.Database.Create(user)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *VerifiedRepository) VerifiedExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Profile{}).Count(&count)
	return count != 0
}


