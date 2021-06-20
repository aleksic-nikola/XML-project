package service

import (
	"fmt"
	"xml/profile-service/data"
	"xml/profile-service/dto"
	"xml/profile-service/repository"
)


type ProfileService struct {
	Repo *repository.ProfileRepository
}

func (service *ProfileService) CreateProfile(profile *data.Profile) error {
	error := service.Repo.CreateProfile(profile)
	return error
}

func (service *ProfileService) GetProfileByUsername(username string) (*data.Profile, error){
	profilObj, error := service.Repo.GetProfileByUsername(username)

	if error!=nil{
		return nil, fmt.Errorf("Can't find any profile obj with username: %s\n", username)
	}
	fmt.Println("IZ SERVISA: ")
	fmt.Println(*profilObj)
	return profilObj, nil
}


func (service *ProfileService) ProfileExists(id uint) (bool, error) {
	
	exists := service.Repo.ProfileExists(id)
	return exists, nil
}

func (service *ProfileService) FollowProfile(myProfile *data.Profile, forFollowProfile *data.Profile) error{
	err := service.Repo.FollowProfile(myProfile, forFollowProfile)
	return err
}

func (service *ProfileService) GetIdByUsername(username string) (uint, error) {
	userId, err := service.Repo.GetIdByUsername(username)

	return userId, err
}

func (service *ProfileService) GetAllFollowingByUsername(username string) []data.Profile {
	followingUsernames := service.Repo.GetAllFollowingByUsername(username)

	return followingUsernames

}

func (service *ProfileService) AcceptFollow(myProfile *data.Profile, newFollower *data.Profile) error {
	err := service.Repo.AcceptFollow(myProfile, newFollower)
	return err
}

func (service *ProfileService) IsUserPublic(username string) (dto.ProfilePublic, error) {
	public, err := service.Repo.IsUserPublic(username)
	return public, err
}

func (service *ProfileService) EditProfileData(dto dto.ProfileEditDTO, oldUsername string) error {
	profile := service.Repo.FindProfileByUsername(oldUsername)

	if profile == nil {
		return nil
	}

	if profile.Phone != dto.Phone {
		if service.Repo.UserExistsByPhone(dto.Phone) {
			return fmt.Errorf("Phone is taken")
		}
	}

	if profile.Username != dto.Username {
		if service.Repo.UserExistsByUsername(dto.Username) {
			return fmt.Errorf("Username is taken")
		}
	}

	// profile update
	profile.Username = dto.Username
	profile.Phone = dto.Phone
	if dto.Gender == 0 {
		profile.Gender = 0
	} else {
		profile.Gender = 1
	}
	profile.DateOfBirth = dto.DateOfBirth
	profile.Website = dto.Website
	profile.Biography = dto.Biography

	err := service.Repo.UpdateProfile(profile)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (service *ProfileService) EditProfilePrivacySettings(privacySettings data.PrivacySetting, username string) error {
	profile := service.Repo.FindProfileByUsername(username)

	if profile == nil {
		return nil
	}

	profile.PrivacySetting.IsPublic = privacySettings.IsPublic
	profile.PrivacySetting.IsInboxOpen = privacySettings.IsInboxOpen
	profile.PrivacySetting.IsTaggingAllowed = privacySettings.IsTaggingAllowed

	err := service.Repo.UpdateProfile(profile)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (service *ProfileService) EditProfileNotificationSettings(notifSettings data.NotificationSetting, username string) error {
	profile := service.Repo.FindProfileByUsername(username)

	if profile == nil {
		return nil
	}

	profile.NotificationSetting.ShowDmNotification = notifSettings.ShowDmNotification
	profile.NotificationSetting.ShowFollowNotification = notifSettings.ShowFollowNotification
	profile.NotificationSetting.ShowTaggedNotification = notifSettings.ShowTaggedNotification

	err := service.Repo.UpdateProfile(profile)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}


func (service *ProfileService) GetCurrentProfile(username string) *data.Profile {
	profile := service.Repo.FindProfileByUsername(username)

	return profile
}

func (service *ProfileService) GetAllFollowersByUsername(username string) []data.Profile   {
	followers := service.Repo.GetAllFollowersByUsername(username)

	return followers
}
