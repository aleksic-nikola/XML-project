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

func (repo *ProfileRepository) GetProfileByUsername(username string) (*data.Profile, error) {

	var profile data.Profile

	result := repo.Database.Where("username = ?", username).First(&profile)

	fmt.Println(profile)

	if result.RowsAffected != 1 {
		error := fmt.Errorf("We didnt find any object with that username!");
		return nil, error
	}

	return &profile, nil
}

func (repo *ProfileRepository) FollowProfile(myProfile *data.Profile, forFollowProfile *data.Profile) error {
	fmt.Println("PRE DODAVANJA MENI FOLLOWING: ")
	fmt.Println(myProfile.Following)

	myProfile.Following = append(myProfile.Following, *forFollowProfile)

	fmt.Println("POSLE DODAVANJA MENI FOLLOWING: ")
	fmt.Println(myProfile.Following)


	forFollowProfile.Followers = append(forFollowProfile.Followers, *myProfile)
	result := repo.Database.Save(myProfile)

	result = repo.Database.Save(forFollowProfile)
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo *ProfileRepository) GetIdByUsername(username string) (uint, error) {
	fullProfile, err := repo.GetProfileByUsername(username)

	if err !=nil{
		return 9999, fmt.Errorf("Greska")
	}

	return fullProfile.ID, err
}

func (repo *ProfileRepository) GetAllFollowingByUsername(username string) []data.Profile {
	var profile data.Profile
	id, _ := repo.GetIdByUsername(username)

	repo.Database.Preload("Following").Find(&profile, id)
	fmt.Println("Ja sam " + profile.Username)
	fmt.Println("******************* JA PRATIM SVE: ********************")
	fmt.Println(profile.Following)

	for _, f := range profile.Following {
		fmt.Println("========================")
		fmt.Println("Ja pratim ---->" + f.Username)

	}
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Broj ljudi koje pratim: ", len(profile.Following))

	return profile.Following
}

func (repo *ProfileRepository) AcceptFollow(myProfile *data.Profile, newFollower *data.Profile) error {

	fmt.Println("JA: REPO:    ", myProfile.Username)
	myProfile.Followers = append(myProfile.Followers, *newFollower)
	result := repo.Database.Save(myProfile)
	fmt.Println(result.RowsAffected)

	if result.RowsAffected==0{
		return fmt.Errorf("Error in repo at accepting profile!")
	}
	fmt.Println("FOLLOWER: REPO:    ", newFollower.Username)

	newFollower.Following = append(newFollower.Following, *myProfile)
	result = repo.Database.Save(newFollower)
	fmt.Println(result.RowsAffected)

	if result.RowsAffected==0{
		return fmt.Errorf("Error in repo at accepting profile!")
	}

	return nil

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
