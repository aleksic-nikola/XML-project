package service

import (
	"xml/campaign-service/data"
	"xml/campaign-service/repository"
)

type MonitoringReportService struct {
	Repo *repository.MonitoringReportRepository
}

func (service *MonitoringReportService) CreateMonitoringReport(monitoringReport *data.MonitoringReport) error {
	error := service.Repo.CreateMonitoringReport(monitoringReport)
	return error
}