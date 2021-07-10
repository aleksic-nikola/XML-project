package service

import (
	"xml/profile-service/data"
	"xml/profile-service/repository")


type AgentService struct {
	Repo *repository.AgentRepository
}

func (service *ProfileService) CreateAgent(profile *data.Profile) error {
	error := service.Repo.CreateProfile(profile)
	return error
}

func (service *ProfileService) AgentExists(id uint) (bool, error) {
	
	exists := service.Repo.ProfileExists(id)
	return exists, nil
}