package service

import (
	"xml/interaction-service/data"
	"xml/interaction-service/repository"
)

type MessageWithContentService struct {
	Repo *repository.MessageWithContentRepository
}

func (service *MessageWithContentService) CreateMessageWithContent(msgWithContent *data.MessageWithContent) error {
	error := service.Repo.CreateMessageWithConent(msgWithContent)
	return error
}
