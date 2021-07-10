package repository

import (
	"fmt"
	"xml/interaction-service/data"

	"gorm.io/gorm"
)

type ProfileNotificationRepository struct {
	Database *gorm.DB
}

func (repo *ProfileNotificationRepository) CreateProfileNotification(profileNotification *data.ProfileNotification) error {
	result := repo.Database.Create(profileNotification)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}
