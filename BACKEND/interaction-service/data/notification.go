package data

import (
	"encoding/json"
	"io"
)

type Notification struct {
	FromUser string `json:"fromUser"`
	ForUser  string `json:"forUser"`
	IsRead   bool   `json:"isread"`
	Text     string `json:"text"`
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
