package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/campaign-service/data"
)

type MultiCampaignRepository struct {
	Database *gorm.DB
}

func (repo *MultiCampaignRepository) CreateMultiCampaign(multiCampaign *data.MultiCampaign) error {
	result := repo.Database.Create(multiCampaign)
	fmt.Println(result.RowsAffected)
	return result.Error
}