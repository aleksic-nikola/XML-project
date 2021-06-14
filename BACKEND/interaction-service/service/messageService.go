package service

import (
	"xml/interaction-service/data"
	"xml/interaction-service/repository"
)

type MessageService struct {
	Repo *repository.MessageRepository
}

func (service *MessageService) CreateMessage(message *data.Message) error {
	error := service.Repo.CreateMessage(message)
	return error
}
