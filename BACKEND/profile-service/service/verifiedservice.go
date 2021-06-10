package service

import (
	"xml/profile-service/data"
	"xml/profile-service/repository")


type VerifiedService struct {
	Repo *repository.VerifiedRepository
}

func (service *VerifiedService) CreateVerified(profile *data.Profile) error {
	error := service.Repo.CreateVerified(profile)
	return error
}

func (service *VerifiedService) VerifiedExists(id uint) (bool, error) {
	
	exists := service.Repo.VerifiedExists(id)
	return exists, nil
}