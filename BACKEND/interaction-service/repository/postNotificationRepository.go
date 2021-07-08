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
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *PostNotificationRepository) GetMyUnreadPostNotif(username string) data.PostNotifications {
	var notifications data.PostNotifications
	repo.Database.Where("is_read = false and for_user = ?", username).Find(&notifications)
	return notifications
}
