package data

import (
	"encoding/json"
	"io"
	"time"

	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	PostedBy  string    `json:"postedby"`
	Timestamp time.Time `json:"timestamp"`
	MediaID   uint      `json:"media_id"`
	Media     Media     `json:"media" gorm:"foreignKey:MediaID"`
	//LocationID            uint      `json:"location_id"`
	//Location              Location  `json:"location" gorm:"foreignKey:LocationID"`
	IsForCloseFriendsOnly bool `json:"isforclosefriendsonly"`
	IsHighlighted         bool `json:"ishighlighted"`
}

func (s *Story) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}

func (s *Stories) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

type Stories []*Story

func GetStories() Stories {
	return storyList
}

var storyList = []*Story{

	{
		//ID:        1,
		PostedBy:  "lucyxz",
		Timestamp: time.Now(),
		Media: Media{
			//MediaID: 1,
			Type: image,
			Path: "some_path",
			Location: Location{
				//LocationID: 1,
				Country: "Serbia",
				City:    "Novi Sad",
			},
		},

		IsForCloseFriendsOnly: true,
		IsHighlighted:         true,
	},
}
