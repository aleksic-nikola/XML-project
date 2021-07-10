package data

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
)

type ProfileNotification struct {
	gorm.Model
	Notification   Notification            `json:"notification" gorm:"embedded"`
	Type           ProfileNotificationType `json:"type" gorm:"type:string"`
}

func (p *ProfileNotification) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(e)
}

type ProfileNotifications []*ProfileNotification

func (p *ProfileNotifications) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProfileNotifications() ProfileNotifications {
	return profileNotificationList
}

var profileNotificationList = []*ProfileNotification{}
	/*{
		Notification: Notification{

			FromUser:      "me",
			For:       "you",
			IsRead:    false,
			Timestamp: time.Now(),
		},
		Type: follow,
	},

	{
		Notification: Notification{

			FromUser:      "me",
			For:       "you",
			IsRead:    true,
			Timestamp: time.Now().AddDate(0, 0, -12),
		},
		Type: message,
	},

	{
		Notification: Notification{

			FromUser:      "me",
			For:       "you",
			IsRead:    true,
			Timestamp: time.Now().AddDate(0, 0, -2),
		},
		Type: message,
	},
}*/
