package service

import (
	"xml/profile-service/data"
	"xml/profile-service/repository")


type VerifiedService struct {
	Repo *repository.VerifiedRepository
}

func (service *VerifiedService) VerifiedExists(id uint) (bool, error) {
	
	exists := service.Repo.VerifiedExists(id)
	return exists, nil
}

func (service *VerifiedService) CreateNewVerified(profile *data.Profile, verifiedType data.VerifiedType) error {

	var verified data.Verified
	verified.Profile = *profile
	verified.Category = verifiedType

	err := service.Repo.CreateVerified(&verified)

	return err
}