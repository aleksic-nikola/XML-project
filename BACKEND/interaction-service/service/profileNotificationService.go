package service

import (
	"xml/interaction-service/data"
	"xml/interaction-service/repository"
)

type ProfileNotificationService struct {
	Repo *repository.ProfileNotificationRepository
}

func (service *ProfileNotificationService) CreateProfileNotification(profileNotificationService *data.ProfileNotification) error {
	error := service.Repo.CreateProfileNotification(profileNotificationService)
	return error
}
