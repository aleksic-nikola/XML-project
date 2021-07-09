package data

import (
	"encoding/json"
	"io"
)

type NotificationSetting struct {
	ShowFollowNotification bool `json:"show_follow_notification"`
	ShowDmNotification     bool `json:"show_dm_notification"`
	ShowTaggedNotification bool `json:"show_tagged_notification"`
}


func (ns *NotificationSetting) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ns)
}

func (ns *NotificationSetting) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ns)
}

func (ns *NotificationSettings) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ns)
}

type NotificationSettings []*NotificationSetting