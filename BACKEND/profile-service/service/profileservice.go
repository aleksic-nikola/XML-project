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

func (service *ProfileService) GetAllFollowersByUsername(username string) []data.Profile {
	followers := service.Repo.GetAllFollowersByUsername(username)

	return followers
}



func (service *ProfileService) BlockProfile(profile_username string, blocked_prof_username string) (*data.Profile,error) {
	profile ,err :=  service.Repo.GetProfileByUsername(profile_username)

	if profile_username == blocked_prof_username {
		fmt.Errorf("You can't block yourself")
		return nil, err
	}

	if err!=nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", profile_username)
		return nil, err
	}

	prof_to_block, error := service.Repo.GetProfileByUsername(blocked_prof_username)

	if error!=nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", blocked_prof_username)
		return nil, error
	}

	profile.Blacklist = append(profile.Blacklist, *prof_to_block)

	err = service.Repo.UpdateProfile(profile)

	if err != nil {
		fmt.Errorf("Error while updating profile - Error in saving blocked user")
		return nil, err
	}

	return profile, nil
}

func (service *ProfileService) MuteProfile(profile_username string, muted_prof_username string) (*data.Profile,error) {
	profile ,err :=  service.Repo.GetProfileByUsername(profile_username)

	if profile_username == muted_prof_username {
		fmt.Errorf("You can't mute yourself")
		return nil, err
	}

	if err!=nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", profile_username)
		return nil, err
	}

	prof_to_mute, error := service.Repo.GetProfileByUsername(muted_prof_username)

	if error!=nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", muted_prof_username)
		return nil, error
	}

	profile.Graylist = append(profile.Graylist, *prof_to_mute)

	err = service.Repo.UpdateProfile(profile)

	if err != nil {
		fmt.Errorf("Error while updating profile - Error in saving muted user")
		return nil, err
	}

	return profile, nil
}

func (service *ProfileService) UnblockProfile(profile_username string, unblocked_prof_username string) (*data.Profile,error) {
	profile ,err :=  service.Repo.GetProfileByUsername(profile_username)

	if profile_username == unblocked_prof_username {
		fmt.Errorf("You can't unblock yourself")
		return nil, err
	}

	if err!=nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", profile_username)
		return nil, err
	}

	prof_to_unblock, error := service.Repo.GetProfileByUsername(unblocked_prof_username)

	if error!=nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", unblocked_prof_username)
		return nil, error
	}

	var newBlackList []data.Profile

	for _,blocked := range profile.Blacklist {
		fmt.Println("BLOKIRAN USER:")
		fmt.Println(blocked.ID)
		if blocked.Username != prof_to_unblock.Username {
			newBlackList = append(newBlackList, blocked)
		}
	}

	err = service.Repo.ClearBlacklist(profile)
	profile.Blacklist = newBlackList
	err = service.Repo.UpdateProfile(profile)

	if err != nil {
		fmt.Errorf("Error while updating profile - Error in saving unblocked user")
		return nil, err
	}

	return profile, nil
}

func (service *ProfileService) UnmuteProfile(profile_username string, unmuted_prof_username string) (*data.Profile,error) {
	profile ,err :=  service.Repo.GetProfileByUsername(profile_username)

	if profile_username == unmuted_prof_username {
		fmt.Errorf("You can't unmute yourself")
		return nil, err
	}

	if err!=nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", profile_username)
		return nil, err
	}

	prof_to_unmute, error := service.Repo.GetProfileByUsername(unmuted_prof_username)

	if error!=nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", unmuted_prof_username)
		return nil, error
	}

	var newGrayList []data.Profile

	for _,muted := range profile.Graylist {
		if muted.Username != prof_to_unmute.Username {
			newGrayList = append(newGrayList, muted)
		}
	}

	err = service.Repo.ClearGrayList(profile)
	profile.Graylist = newGrayList
	err = service.Repo.UpdateProfile(profile)

	if err != nil {
		fmt.Errorf("Error while updating profile - Error in saving unmuted user")
		return nil, err
	}

	return profile, nil
}
