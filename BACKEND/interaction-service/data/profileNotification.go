package data

import (
	"encoding/json"
	"io"
	"time"
)

type ProfileNotification struct {
	Notification
	Type ProfileNotificationType `json:"type"`
}

func(p *ProfileNotification) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(e)
}

type ProfileNotifications []*ProfileNotification

func(p *ProfileNotifications) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProfileNotifications() ProfileNotifications {
	return profileNotificationList
}

var profileNotificationList = []*ProfileNotification {
	{
		Notification : Notification{
			ID: 1,
			From : "me",
			For : "you",
			IsRead: false,
			Timestamp: time.Now(),
		},
		Type: follow,
	},

	{
		Notification : Notification{
			ID: 2,
			From : "me",
			For : "you",
			IsRead: true,
			Timestamp: time.Now().AddDate(0, 0, -12),
		},
		Type: message,
	},

	{
		Notification : Notification{
			ID: 3,
			From : "me",
			For : "you",
			IsRead: true,
			Timestamp: time.Now().AddDate(0, 0, -2),
		},
		Type: message,
	},

}
