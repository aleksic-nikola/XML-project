package service

import (
	"xml/request-service/data"
	"xml/request-service/repository"
)

type MonitorReportRequestService struct {
	Repo *repository.MonitorReportRequestRepository
}

func (service *MonitorReportRequestService) CreateMonitorReportRequest(monitorReportRequest *data.MonitorReportRequest) error {
	error := service.Repo.CreateMonitorReportRequest(monitorReportRequest)
	return error
}
