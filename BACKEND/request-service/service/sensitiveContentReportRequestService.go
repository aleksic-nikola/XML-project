package service

import (
	"xml/request-service/data"
	"xml/request-service/repository"
)

type SensitiveContentReportRequestService struct {
	Repo *repository.SensitiveContentReportRequestRepository
}

func (service *SensitiveContentReportRequestService) CreateSensitiveContentReportRequest(sensitiveContentReportRequest *data.SensitiveContentReportRequest) error {
	error := service.Repo.CreateSensitiveContentReportRequest(sensitiveContentReportRequest)
	return error
}