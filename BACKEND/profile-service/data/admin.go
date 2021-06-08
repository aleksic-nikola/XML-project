package data

import (
	"encoding/json"
	"io"
	"time"
)

type Admin struct {
	Profile Profile `json:"profile"`
}

func (a *Admin) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func (a* Admins) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

type Admins []*Admin

func GetAdmins() Admins {
	return adminList
}

var adminList = []*Admin{
	{
		Profile: Profile{
			ID: 1,
			Name: "Danilo",
			Lastname: "Paripovic",
			Email: "danilo@gmail.com",
			Phone: "03214321",
			Gender: MALE,
			DateOfBirth: time.Date(1998, time.September, 29, 0, 0, 0, 0, time.UTC),
			Website: "www.danilo.com",
			Biography: "prazna",
			CloseFriends: []string{"mark","nikolat","nikolaa"},
			Favourites: nil,
			IsBanned: false,
			PrivacySetting: PrivacySetting{
				IsTaggingAllowed: true,
				IsPublic: true,
				IsInboxOpen: true,
				Graylist: nil,
				Blacklist: nil,
			},
			NotificationSetting: NotificationSetting {
				ShowDmNotification: true,
				ShowFollowNotification: true,
				ShowTaggedNotification: true,
			},
		},
	},
}
