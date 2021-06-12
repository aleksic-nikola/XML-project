package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/campaign-service/data"
)

type MonitoringReportRepository struct {
	Database *gorm.DB
}

func (repo *MonitoringReportRepository) CreateMonitoringReport(monitoringReport *data.MonitoringReport) error {
	result := repo.Database.Create(monitoringReport)
	fmt.Println(result.RowsAffected)
	return result.Error
}