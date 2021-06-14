package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/request-service/data"
)

type SensitiveContentReportRequestRepository struct {
	Database *gorm.DB
}

func (repo *SensitiveContentReportRequestRepository) CreateSensitiveContentReportRequest(sensitiveContentReportRequest *data.SensitiveContentReportRequest) error {
	result := repo.Database.Create(sensitiveContentReportRequest)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}