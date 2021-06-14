package data

import (
	"encoding/json"
	"io"
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	From      string    `json:"from"`
	For       string    `json:"for"`
	IsRead    bool      `json:"isread"`
	Timestamp time.Time `json:"timestamp"`
}

func (m *Notification) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(e)
}

// collection
type Notifications []*Notification

func (n *Notifications) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(n)
}

func GetNotifications() Notifications {
	return notificationList
}

var notificationList = []*Notification{}
