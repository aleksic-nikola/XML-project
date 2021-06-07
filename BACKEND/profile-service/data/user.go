package data

import (
	"encoding/json"
	"io"
	"time"
)

type Profile struct {
	ID              int                 	`json:"id"`
	Name            string              	`json:"name"`
	Lastname        string              	`json:"lastname"`
	Email           string              	`json:"email"`
	Phone           string              	`json:"phone"`
	Gender          Gender              	`json:"gender"`
	DateOfBirth     time.Time          		`json:"date_of_birth"`
	Website         string         	        `json:"website"`
	Biography       string       	        `json:"biography"`
	CloseFriends    []string 	            `json:"close_friends"`
	Favourites      map[string][]string		`json:"favourites"`
	IsBanned        bool                	`json:"is_banned"`
	PrivacySetting  PrivacySetting          `json:"privacy_settings"`
	NotificationSetting NotificationSetting `json:"notification_setting"`
}

func (u *Profile) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u*Profiles) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

type Profiles []*Profile

func GetProfiles() Profiles {
	return profileList
}

var profileList = []*Profile{

	{
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
}