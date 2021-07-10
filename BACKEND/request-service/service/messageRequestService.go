package service

import (
	"xml/request-service/data"
	"xml/request-service/repository"
)

type MessageRequestService struct {
	Repo *repository.MessageRequestRepository
}

func (service *MessageRequestService) CreateMessageRequest(messageRequest *data.MessageRequest) error {
	error := service.Repo.CreateMessageRequest(messageRequest)
	return error
}