package repository

import (
	"fmt"
	"xml/content-service/data"

	"gorm.io/gorm"
)

type MediaRepository struct {
	Database *gorm.DB
}

func (repo *MediaRepository) CreateMedia(media *data.Media) error {
	result := repo.Database.Create(media)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *MediaRepository) MediaExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Media{}).Count(&count)
	return count != 0
}
