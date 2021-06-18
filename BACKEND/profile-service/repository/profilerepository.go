package repository

import (
	"fmt"
	"xml/profile-service/data"
	"xml/profile-service/dto"

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

func (repo *ProfileRepository) IsUserPublic(username string) (dto.ProfilePublic, error) {
	if !repo.UserExistsByUsername(username) {
		return dto.ProfilePublic{}, fmt.Errorf("no user with that username")
	}
	var profile data.Profile
	repo.Database.Where("username = ?", username).Find(&data.Profile{}).First(&profile)
	var dto = dto.ProfilePublic{Public: profile.PrivacySetting.IsPublic}

	return dto, nil
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

func(repo *ProfileRepository) UserExistsByUsername(username string) bool {

	var count int64
	repo.Database.Where("username = ?", username).Find(&data.Profile{}).Count(&count)
	return count != 0
}