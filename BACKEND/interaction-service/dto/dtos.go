package dto

import (
	"encoding/json"
	"io"
)

type PostNotif struct {
	Post_id int `json:"post_id"`
	Type string `json:"type"`
}

func (pn *PostNotif) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(pn)
}

func (pn *PostNotif) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(pn)
}

type UsernameRole struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (u *UsernameRole) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *UsernameRole) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

type UsernameDto struct {
	Username string `json:"username"`
}

func (ur *UsernameDto) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ur)
}

func (ur *UsernameDto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ur)
}

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