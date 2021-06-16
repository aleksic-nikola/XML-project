package repository

import (
	"fmt"
	"xml/profile-service/data"
	"xml/profile-service/dtos"

	"gorm.io/gorm"
)
type ProfileRepository struct {
	Database *gorm.DB
}

func (repo *ProfileRepository) CreateProfile(user *data.Profile) error {
	result := repo.Database.Create(user)
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *ProfileRepository) ProfileExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Profile{}).Count(&count)
	return count != 0
}

func (repo *ProfileRepository) IsUserPublic(username string) (dtos.ProfilePublic, error) {
	if !repo.ProfileExistsByUsername(username) {
		return dtos.ProfilePublic{}, fmt.Errorf("no user with that username")
	}
	var profile data.Profile
	repo.Database.Where("username = ?", username).Find(&data.Profile{}).First(&profile)
	var dto = dtos.ProfilePublic{Public: profile.PrivacySetting.IsPublic}

	return dto, nil
}

func (repo *ProfileRepository) ProfileExistsByUsername(username string) bool {
	var count int64
	repo.Database.Where("username = ?", username).Find(&data.Profile{}).Count(&count)
	return count != 0
}