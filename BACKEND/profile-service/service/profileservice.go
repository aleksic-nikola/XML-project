package service

import (
	"fmt"
	"time"
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
	layoutISO := "2006-01-02"
	//date := "1999-12-31"
	t, _ := time.Parse(layoutISO, dto.DateOfBirth)

	profile.DateOfBirth = t
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

func (service *ProfileService) AddPostToFavourites(roleDto dto.UsernameRole, favourites dto.PostToFavourites) error {

	user, err := service.Repo.GetProfileByUsername(roleDto.Username)

	if err != nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", roleDto.Username)
		return err
	}

	fmt.Println("USER:")
	fmt.Println(user)
	fmt.Println("FAVORITI:")
	fmt.Println(user.Favourites)

	var newSave data.SavedPost
	newSave.PostId = favourites.PostId

	var collectionExist = false
	for i,collection := range user.Favourites {
		if collection.CollectionName == favourites.CollectionName {
			fmt.Println("USAO OVDE")
			for _, j := range collection.SavedPosts {
				if j.PostId == newSave.PostId {
					err1 := fmt.Errorf("POST ALREADY EXISTS IN THIS COLLECTION");
					return err1
				}
			}
			user.Favourites[i].SavedPosts = append(collection.SavedPosts, newSave)
			fmt.Println(collection.SavedPosts)
			collectionExist = true
		}
	}

	if collectionExist == false {
		var newFavourite data.Favourite
		newFavourite.CollectionName = favourites.CollectionName
		newFavourite.SavedPosts = append(newFavourite.SavedPosts, newSave)

		user.Favourites = append(user.Favourites, newFavourite)
	}

	fmt.Println(user)
	err = service.Repo.UpdateProfile(user)

	if err != nil {
		fmt.Errorf("Error while updating profile - Error in saving unmuted user")
		return err
	}

	return nil
}

func (service *ProfileService) GetPostsIdsInCollection(collection string, user dto.UsernameRole) (dto.PostIdsDto,error) {
	profile, err := service.Repo.GetProfileByUsername(user.Username)

	if err != nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", user.Username)
		return dto.PostIdsDto{}, err
	}

	var posts dto.PostIdsDto

	for _,f := range profile.Favourites {
		if f.CollectionName == collection {
			for _,p := range f.SavedPosts {
				fmt.Println("POSTOVI::::")
				fmt.Println(p)
				var post dto.PostIdDto
				post.Id = p.PostId
				posts.Ids = append(posts.Ids, post)
			}
		}
	}

	fmt.Println("**********************************")

	for _, p := range posts.Ids {
		fmt.Println("POST:")
		fmt.Println(p.Id)
	}
	fmt.Println("**********************************")


	return posts, nil
}

func (service *ProfileService) GetAllPublicProfiles() (error, data.Profiles) {

	err, profiles := service.Repo.GetAllPublicProfiles()
	return err, profiles
}

// gets all profiles the user follows but are not public
func (service *ProfileService) GetAllAllowedProfiles(username string) (data.Profiles) {
	profiles := service.Repo.GetAllFollowingByUsername(username)
	var filtered_profiles data.Profiles
	for _, p := range profiles {
		fmt.Println("POST:")
		if p.PrivacySetting.IsPublic == true {
			continue
		}
		// private profile
		filtered_profiles = append(filtered_profiles, &p)
	}

	return filtered_profiles
}

func (service *ProfileService) GetAllNonFollowedPrivateProfiles(username string) (error, data.Profiles){

	following := service.Repo.GetAllFollowingByUsername(username)
	err, private := service.Repo.GetAllPrivateProfiles(username)
	var retList data.Profiles
	for _, p := range private {
		isFollowing := false
		for _, f := range following {
			if f.Username == p.Username {
				isFollowing = true
				continue
			}
		}
		if isFollowing {
			continue
		}
		// not following add it
		retList = append(retList, p)
	}

	return err, retList
}

func (service *ProfileService) GetFavouritPostsIds(roleDto dto.UsernameRole) (dto.Collection, error) {
	profile, err := service.Repo.GetProfileByUsername(roleDto.Username)

	if err != nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", roleDto.Username)
		return dto.Collection{}, err
	}

	var collections dto.Collection

	for _, f := range profile.Favourites {
		collections.Name = append(collections.Name, f.CollectionName)
	}

	for _, cn := range collections.Name {
		fmt.Println(cn)
	}

	return collections, nil
}

func (service *ProfileService) DeleteCollection(username string, collectionName string) error {
	profile, err := service.Repo.GetProfileByUsername(username)

	if err != nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", username)
		return err
	}

	var newFavourites []data.Favourite

	for _,col := range profile.Favourites {
		if col.CollectionName != collectionName {
			newFavourites = append(newFavourites, col)
		}
	}

	err = service.Repo.ClearFavourites(profile)
	profile.Favourites = newFavourites
	err = service.Repo.UpdateProfile(profile)

	if err != nil {
		fmt.Errorf("Error while updating profile - Error in deleting collection")
		return err
	}

	return nil
}

func (service *ProfileService) DeletePostFromCollection(username string, favourites dto.PostToFavourites) error {
	fmt.Println(username)
	profile, err := service.Repo.GetProfileByUsername(username)

	if err != nil{
		fmt.Errorf("Can't find any profile obj with username: %s\n", username)
		return err
	}

	var newFavourites []data.Favourite
	var singleFav data.Favourite

	for _,col := range profile.Favourites {
		if col.CollectionName == favourites.CollectionName {
			var newSavedPostList []data.SavedPost
			for _, post := range col.SavedPosts {
				if post.PostId != favourites.PostId {
					newSavedPostList = append(newSavedPostList, post)
				}
			}
			singleFav.SavedPosts = newSavedPostList
			singleFav.CollectionName = col.CollectionName
			newFavourites = append(newFavourites, singleFav)

		} else {
			newFavourites = append(newFavourites, col)
		}
	}

	err = service.Repo.ClearFavourites(profile)
	profile.Favourites = newFavourites
	err = service.Repo.UpdateProfile(profile)

	if err != nil {
		fmt.Errorf("Error while updating profile - Error in deleting collection")
		return err
	}

	return nil
}

func (service *ProfileService) GetUsersWhoBlockedMe(myusername string) (error, []string) {

	var listOfUsersWhoBlockedMe []string
	var listOfAllProfiles data.Profiles

	err, listOfAllProfiles :=  service.Repo.GetAllProfiles()

	fmt.Println("============================")
	fmt.Println(len(listOfAllProfiles))

	if err != nil{
		fmt.Errorf("Can't get allprofiles")
		return err, nil
	}

	for _,p := range listOfAllProfiles {

		for _,b := range p.Blacklist {
			if b.Username == myusername {
				listOfUsersWhoBlockedMe = append(listOfUsersWhoBlockedMe, p.Username)
			}
		}
	}

	return err, listOfUsersWhoBlockedMe
}

func (service *ProfileService) GetMyNotificationsSettings(username string) (data.NotificationSetting,error) {
	profile, err := service.Repo.GetProfileByUsername(username)

	if err != nil {
		return data.NotificationSetting{}, err
	}

	notificationSetting := profile.NotificationSetting

	return notificationSetting, nil
}


func (service *ProfileService) GetCloseFriendsByUsername(username string) ([]data.Profile, error) {
	closeFriends, err := service.Repo.GetCloseFriendsByUsername(username)

	return closeFriends, err
}

func (service *ProfileService) AddProfileToCloseFriends(myUsername string, usernameForAddToCloseFriends string) error {

	err := service.Repo.AddProfileToCloseFriends(myUsername, usernameForAddToCloseFriends)

	return err

}

func (service *ProfileService) RemoveProfileFromCloseFriends(myUsername string, usernameForRemoveFromCloseFriends string) error {
	err := service.Repo.RemoveProfileFromCloseFriends(myUsername, usernameForRemoveFromCloseFriends)

	return err
}
