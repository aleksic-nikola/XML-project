package data

import (
	"encoding/json"
	"io"
	//"time"

)

type Agent struct {
	ProfileID   string  `json:"profile_id"`
	Profile    Profile `json:"profile" gorm:"foreignkey:ProfileID"`
	Webshop    string  `json:"webshop"`
}

func (a *Agent) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func (a* Agents) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

type Agents []*Agent

func GetAgents() Agents {
	return agentList
}

var agentList = []*Agent{}
/*
var agentList = []*Agent{
	{
		Profile: Profile{
			Username : "dparip",
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
		Webshop: "nekisajt.com",
	},
}
*/