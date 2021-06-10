package data

import (
	"encoding/json"
	"io"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Username        string              	`json:"username" gorm:"uniqueIndex"`
	Phone           string              	`json:"phone" gorm:"uniqueIndex"`
	Gender          Gender              	`json:"gender" gorm:"type:int" `
	DateOfBirth     time.Time          	`json:"date_of_birth;type:date"`
	Website         string         	        `json:"website"`
	Biography       string       	        `json:"biography"`
	CloseFriends    []Profile 	        `json:"close_friends" gorm:"many2many:profile_close_friends;"`
	Favourites      map[string][]string	`json:"favourites" gorm:"type:text"`
	IsBanned        bool                	`json:"is_banned"`
	PrivacySetting  PrivacySetting          `gorm:"embedded"`
	NotificationSetting NotificationSetting `json:"notification_setting" gorm:"embedded"`
	Following 	[]Profile 		`json:"following" gorm:"many2many:profile_following;"`
	Followers       []Profile 		`json:"followers" gorm:"many2many:profile_followers;"`	
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

var profileList = []*Profile {}
/*
var profileList = []*Profile{

	{
		Username: "dparip",
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
*/