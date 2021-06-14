package service

import (
	"xml/request-service/data"
	"xml/request-service/repository"
)

type VerificationRequestService struct {
	Repo *repository.VerificationRequestRepository
}

func (service *VerificationRequestService) CreateVerificationRequest(verificationRequest *data.VerificationRequest) error {
	error := service.Repo.CreateVerificationRequest(verificationRequest)
	return error
}

