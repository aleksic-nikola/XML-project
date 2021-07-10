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

	result := repo.Database.Preload("Blacklist").Preload("Graylist").Preload("Following").
		Preload("Followers").Preload("CloseFriends").Preload("Favourites.SavedPosts").Where("username = ?", username).First(&profile)

	fmt.Println(profile)
	fmt.Println("************************************************")
	fmt.Println(username)


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

	repo.Database.Preload("Following").Preload("Graylist").Preload("Blacklist").Find(&profile, id)
	//fmt.Println("Ja sam " + profile.Username)
	//fmt.Println("******************* JA PRATIM SVE: ********************")
	//fmt.Println(profile.Following)

	var filtered []data.Profile
	for _, f := range profile.Following {
		//fmt.Println("========================")
		//fmt.Println("Ja pratim ---->" + f.Username)

		prof, err :=  repo.GetProfileByUsername(f.Username)

		if err != nil {
			fmt.Errorf("Error in GetProfileByUsername")
		}

		flag := false
		for _, b := range prof.Blacklist {
			if username == b.Username {
				flag = true
				break
			}
		}

		if flag == false {
			filtered = append(filtered, f)
		}

	}

	//fmt.Println("---------------------------------------------------------")
	//fmt.Println("Broj ljudi koje pratim: ", len(profile.Following))

	return filtered
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
	repo.Database.Preload("Graylist").Preload("Blacklist").Where("username =?", username).First(&profile)
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


func (repo *ProfileRepository) GetAllFollowersByUsername(username string) []data.Profile {
	var profile data.Profile
	id, _ := repo.GetIdByUsername(username)

	repo.Database.Preload("Followers").Find(&profile, id)

	return profile.Followers
}

func (repo *ProfileRepository) GetProfileByID(my_id uint) (*data.Profile, error) {

	var profile data.Profile

	result := repo.Database.Where("id = ?", my_id).First(&profile)

	fmt.Println(profile)

	if result.RowsAffected != 1 {
		error := fmt.Errorf("We didnt find any object with that username!");
		return nil, error
	}

	return &profile, nil
}

func (repo *ProfileRepository) ClearBlacklist(profile *data.Profile) error {

	repo.Database.Model(&profile).Association("Blacklist").Clear()

	return nil
}

func (repo *ProfileRepository) ClearGrayList(profile *data.Profile) error {

	repo.Database.Model(&profile).Association("Graylist").Clear()

	return nil
}

func (repo *ProfileRepository) GetAllPublicProfiles() (error, data.Profiles) {

	var profiles data.Profiles
	err := repo.Database.Preload("Blacklist").Preload("Graylist").Where("is_public = ?", true).Find(&profiles)

	if err.Error != nil {
		fmt.Println(err.Error)
		return fmt.Errorf("there has been an error retrieving public profiles"), nil
	}
	return nil, profiles
}

func (repo *ProfileRepository) GetAllPrivateProfiles(username string) (error, data.Profiles) {

	var profiles data.Profiles

	err := repo.Database.Preload("Graylist").Preload("Blacklist").Where("is_public = ? AND username != ?", false, username).Find(&profiles)

	if err.Error != nil {
		fmt.Println(err.Error)
		return fmt.Errorf("there has been an error retrieving public profiles"), nil
	}
	return nil, profiles
}

func (repo *ProfileRepository) ClearFavourites(profile *data.Profile) error {
	repo.Database.Model(&profile).Association("Favourites").Clear()

	return nil
}

func (repo *ProfileRepository) GetAllProfiles() (error, data.Profiles) {

	var profiles data.Profiles

	err := repo.Database.Preload("Blacklist").Preload("Graylist").Find(&profiles)

	if err.Error != nil {
		fmt.Println(err.Error)
		return fmt.Errorf("there has been an error retrieving profiles"), nil
	}
	return nil, profiles
}

func (repo *ProfileRepository) GetCloseFriendsByUsername(username string) ([]data.Profile, error) {
	var userProfile data.Profile
	res := repo.Database.Preload("CloseFriends").Where("username = ?", username).Find(&userProfile)

	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, fmt.Errorf("there has been an error retrieving CloseFriends")

	}

	if res.RowsAffected == 0{
		fmt.Println(res.Error)
		return nil, fmt.Errorf("User not found!")
	}

	fmt.Println(userProfile.CloseFriends)
	return userProfile.CloseFriends, nil

}

func (repo *ProfileRepository) AddProfileToCloseFriends(myUsername string, usernameForAddToCloseFriends string) error {

	var myProfile data.Profile
	res := repo.Database.Preload("CloseFriends").Where("username = ?", myUsername).Find(&myProfile)

	if res.Error != nil {
		fmt.Println(res.Error)
		return fmt.Errorf("there has been an error retrieving myProfile")

	}
	if res.RowsAffected == 0{
		fmt.Println(res.Error)
		return fmt.Errorf("MyProfile not found")
	}

	var profileForAddToCloseFriend data.Profile
	res = repo.Database.Where("username = ?", usernameForAddToCloseFriends).Find(&profileForAddToCloseFriend)

	if res.Error != nil {
		fmt.Println(res.Error)
		return fmt.Errorf("there has been an error retrieving profileForAddToCloseFriend")

	}
	if res.RowsAffected == 0{
		fmt.Println(res.Error)
		return fmt.Errorf("profileForAddToCloseFriend not found")
	}

	fmt.Println("PROFIL KOJI DODAJEM: ", profileForAddToCloseFriend.Username)
	fmt.Println("JA SAM : ", myProfile.Username)

	isAlreadyInList, _ := containsInCloseFriends(myProfile.CloseFriends, profileForAddToCloseFriend.Username)

	if isAlreadyInList{
		return fmt.Errorf("profile is already in CloseFriends")
	}


	myProfile.CloseFriends = append(myProfile.CloseFriends, profileForAddToCloseFriend)

	res = repo.Database.Save(myProfile)

	if res.Error != nil {
		fmt.Println(res.Error)
		return fmt.Errorf("there has been an error adding to closeFriends")
	}

	return nil


}

func (repo *ProfileRepository) RemoveProfileFromCloseFriends(myUsername string, usernameForRemoveFromCloseFriends string) error{
	var myProfile data.Profile
	res := repo.Database.Preload("CloseFriends").Where("username = ?", myUsername).Find(&myProfile)

	if res.Error != nil {
		fmt.Println(res.Error)
		return fmt.Errorf("there has been an error retrieving myProfile")

	}
	if res.RowsAffected == 0{
		fmt.Println(res.Error)
		return fmt.Errorf("MyProfile not found")
	}

	isInList, _ := containsInCloseFriends(myProfile.CloseFriends, usernameForRemoveFromCloseFriends)
	if !isInList {
		return fmt.Errorf("Profile is not in closeFriends list")
	}

	fmt.Println("PRE BRISANJA: ")
	for _, oneFriend := range myProfile.CloseFriends{
		fmt.Println(oneFriend.Username)
	}
	fmt.Println("***************************************")


	var newCloseFriendsList []data.Profile

	for _,oneProfile := range myProfile.CloseFriends {
		if oneProfile.Username != usernameForRemoveFromCloseFriends {
			newCloseFriendsList = append(newCloseFriendsList, oneProfile)
		}
	}


	errClear:= repo.Database.Model(&myProfile).Association("CloseFriends").Clear()
	if errClear != nil {
		fmt.Println(res.Error)
		return fmt.Errorf("Error clearing table")
	}

	myProfile.CloseFriends = newCloseFriendsList

	res = repo.Database.Save(myProfile)

	if res.Error != nil {
		fmt.Println(res.Error)
		return fmt.Errorf("there has been an error removing from closeFriends")
	}

	return nil
}

func (repo *ProfileRepository) ClearFollowing(profile *data.Profile) error {

	return repo.Database.Model(&profile).Association("Following").Clear()

}

func (repo *ProfileRepository) ClearFollowers(profile *data.Profile) error {

	return repo.Database.Model(&profile).Association("Followers").Clear()

}

func removeFromList(arr []data.Profile, index int) []data.Profile {
	return append(arr[:index], arr[index+1:]...)
}

func containsInCloseFriends(closeFriends []data.Profile, username string) (bool, int) {

	for i, oneProfile:= range closeFriends{
		if oneProfile.Username == username {
			return true, i
		}
	}
	return false, -1

}