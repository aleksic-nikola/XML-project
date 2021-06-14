package repository

import (
	"fmt"
	"xml/content-service/data"

	"gorm.io/gorm"
)

type StoryRepository struct {
	Database *gorm.DB
}

func (repo *StoryRepository) CreateStory(story *data.Story) error {
	result := repo.Database.Create(story)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *StoryRepository) StoryExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Story{}).Count(&count)
	return count != 0
}
