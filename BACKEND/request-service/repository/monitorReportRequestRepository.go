package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/request-service/data"
)

type MonitorReportRequestRepository struct {
	Database *gorm.DB
}

func (repo *MonitorReportRequestRepository) CreateMonitorReportRequest(monitorReportRequest *data.MonitorReportRequest) error {
	result := repo.Database.Create(monitorReportRequest)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

