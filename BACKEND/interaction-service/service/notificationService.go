package service

import (
	"xml/interaction-service/data"
	"xml/interaction-service/repository"
)

type NotificationService struct {
	Repo *repository.NotificationRepository
}

func (service *NotificationService) CreateNotification(notification *data.Notification) error {
	error := service.Repo.CreateNotification(notification)
	return error
}
