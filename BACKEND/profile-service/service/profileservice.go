package service

import (
	"fmt"
	"xml/profile-service/data"
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