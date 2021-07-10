package service

import (
	"xml/interaction-service/data"
	"xml/interaction-service/dto"
	"xml/interaction-service/repository"
)

type PostNotificationService struct {
	Repo *repository.PostNotificationRepository
}

func (s PostNotificationService) CreatePostNotification(notificationDto dto.PostNotif, notificationFrom string, notificationFor string) error {
	var postnotif data.PostNotification

	postnotif.PostID = notificationDto.Post_id
	postnotif.Type = data.PostNotificationType(notificationDto.Type)
	postnotif.Notification.IsRead = false
	postnotif.Notification.FromUser = notificationFrom
	postnotif.Notification.ForUser = notificationFor

	err := s.Repo.CreatePostNotification(&postnotif)

	return err
}

func (s PostNotificationService) GetMyUnreadPostNotif(username string) (data.PostNotifications) {
	notifications := s.Repo.GetMyUnreadPostNotif(username)
	return notifications
}

func (s PostNotificationService) ReadPostNotifications(username string) error {
	notifications := s.Repo.GetMyUnreadPostNotif(username)

	for _,n := range notifications {
		n.Notification.IsRead = true
		err := s.Repo.Save(n)
		if err != nil {
			return err
		}
	}
	return nil
}


