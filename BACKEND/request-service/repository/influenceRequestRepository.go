package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/request-service/data"
)

type InfluenceRequestRepository struct {
	Database *gorm.DB
}

func (repo *InfluenceRequestRepository) CreateInfluenceRequest(influenceRequest *data.InfluenceRequest) error {
	result := repo.Database.Create(influenceRequest)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}