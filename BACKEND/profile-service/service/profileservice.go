package service

import (
	"xml/profile-service/data"
	"xml/profile-service/dtos"
	"xml/profile-service/repository")


type ProfileService struct {
	Repo *repository.ProfileRepository
}

func (service *ProfileService) CreateProfile(profile *data.Profile) error {
	error := service.Repo.CreateProfile(profile)
	return error
}

func (service *ProfileService) ProfileExists(id uint) (bool, error) {
	
	exists := service.Repo.ProfileExists(id)
	return exists, nil
}

func (service *ProfileService) IsUserPublic(username string) (dtos.ProfilePublic, error) {
	public, err := service.Repo.IsUserPublic(username)
	return public, err
}