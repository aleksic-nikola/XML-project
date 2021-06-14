package service

import (
	"xml/interaction-service/data"
	"xml/interaction-service/repository"
)

type MessageWithOneTimeContentService struct {
	Repo *repository.MessageWithOneTimeContentRepository
}

func (service *MessageWithOneTimeContentService) CreateMessageWithOneTimeContent(msgWithOneTimeContent *data.MessageWithOneTimeContent) error {
	error := service.Repo.CreateMessageWithOneTimeConent(msgWithOneTimeContent)
	return error
}
