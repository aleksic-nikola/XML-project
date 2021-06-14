package data

import (
	"encoding/json"
	"io"
	"time"
)

type PostNotification struct {
	NotificationID uint                 `json:"notification_id"`
	Notification   Notification         `json:"notification" gorm:"foreignkey=NotificationID"`
	PostID         int                  `json:"postid"`
	Type           PostNotificationType `json:"type" gorm:"type:string"`
}

func (p *PostNotification) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(e)
}

type PostNotifications []*PostNotification

func (p *PostNotifications) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetPostNotifications() PostNotifications {
	return postNotificationList
}

var postNotificationList = []*PostNotification{
	{
		Notification: Notification{

			From:      "me",
			For:       "you",
			IsRead:    false,
			Timestamp: time.Now(),
		},
		PostID: 12,
		Type:   comment,
	},

	{
		Notification: Notification{

			From:      "me",
			For:       "you",
			IsRead:    true,
			Timestamp: time.Now().AddDate(0, 0, -12),
		},
		PostID: 12,
		Type:   tag,
	},

	{
		Notification: Notification{

			From:      "me",
			For:       "you",
			IsRead:    true,
			Timestamp: time.Now().AddDate(0, 0, -2),
		},
		PostID: 12,
		Type:   like,
	},
}
