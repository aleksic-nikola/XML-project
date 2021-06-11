package repository

import (
	"fmt"
	"xml/content-service/data"

	"gorm.io/gorm"
)

type LocationRepository struct {
	Database *gorm.DB
}

func (repo *LocationRepository) CreateLocation(location *data.Location) error {
	result := repo.Database.Create(location)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *LocationRepository) LocationExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Location{}).Count(&count)
	return count != 0
}
