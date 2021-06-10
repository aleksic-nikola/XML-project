package repository

import (
	"fmt"
	"xml/interaction-service/data"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	Database *gorm.DB
}

func (repo *NotificationRepository) CreateNotification(notification *data.Notification) error {
	result := repo.Database.Create(notification)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}
