package service

import (
	"xml/interaction-service/data"
	"xml/interaction-service/repository"
)

type PostNotificationService struct {
	Repo *repository.PostNotificationRepository
}

func (service *PostNotificationService) CreatePostNotification(postNotification *data.PostNotification) error {
	error := service.Repo.CreatePostNotification(postNotification)
	return error
}
