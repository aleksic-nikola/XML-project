package repository

import (
	"fmt"
	"xml/profile-service/data"

	"gorm.io/gorm"
)
type ProfileRepository struct {
	Database *gorm.DB
}

func (repo *ProfileRepository) CreateProfile(user *data.Profile) error {
	result := repo.Database.Create(user)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *ProfileRepository) ProfileExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Profile{}).Count(&count)
	return count != 0
}

func(repo *ProfileRepository) FindProfileByUsername(username string) *data.Profile {
	var profile data.Profile
	repo.Database.Where("username =?", username).First(&profile)
	return &profile
}

func(repo *ProfileRepository) UserExistsByPhone(phone string) bool {
	var count int64
	repo.Database.Where("phone = ?", phone).Find(&data.Profile{}).Count(&count)
	return count != 0
}

func(repo *ProfileRepository) UpdateProfile(profile *data.Profile) error {
	err := repo.Database.Save(&profile).Error

	return err
}