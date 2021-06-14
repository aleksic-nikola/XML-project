package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/campaign-service/data"
)

type OneTimeCampaignRepository struct {
	Database *gorm.DB
}

func (repo *OneTimeCampaignRepository) CreateOneTimeCampaign(oneTimeCampaign *data.OneTimeCampaign) error {
	result := repo.Database.Create(oneTimeCampaign)
	fmt.Println(result.RowsAffected)
	return result.Error
}
