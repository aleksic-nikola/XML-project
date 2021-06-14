package service

import (
	"xml/request-service/data"
	"xml/request-service/repository"
)

type InfluenceRequestService struct {
	Repo *repository.InfluenceRequestRepository
}

func (service *InfluenceRequestService) CreateInfluenceRequest(influenceRequest *data.InfluenceRequest) error {
	error := service.Repo.CreateInfluenceRequest(influenceRequest)
	return error
}

