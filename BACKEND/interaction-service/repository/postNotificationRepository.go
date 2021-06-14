package repository

import (
	"fmt"
	"xml/interaction-service/data"

	"gorm.io/gorm"
)

type PostNotificationRepository struct {
	Database *gorm.DB
}

func (repo *PostNotificationRepository) CreatePostNotification(postNotification *data.PostNotification) error {
	result := repo.Database.Create(postNotification)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}
