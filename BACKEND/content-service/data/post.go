package data

import (
	"encoding/json"
	"io"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	PostedBy    string    `json:"postedby"`
	Timestamp   time.Time `json:"timestamp"`
	Description string    `json:"description"`
	Comments    []Comment `json:"comments"`
	//LocationID  uint      `json:"location_id"`
	//Location    Location  `json:"location" gorm:"foreignKey:LocationID" `
	Medias   []Media  `json:"medias"`
	Likes    []string `json:"likes" gorm:"type:text"`
	Dislikes []string `json:"dislikes" gorm:"type:text"`
}

func (p *Post) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Posts) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

type Posts []*Post

func GetPosts() Posts {
	return postList
}

var postList = []*Post{

	{
		//ID:          1,
		PostedBy:    "lucyxz",
		Timestamp:   time.Now(),
		Description: "the description",
		Comments: []Comment{
			{
				//CommentID: 1,
				PostedBy:  "lucyxz",
				Text:      "some text here",
				Timestamp: time.Now(),
			},
		},

		Medias: []Media{
			{
				//MediaID: 1,
				Type: image,
				Path: "some_path",
				Location: Location{
					//LocationID: 1,
					Country: "Serbia",
					City:    "Novi Sad",
				},
			},
		},
		Likes:    []string{"like1"},
		Dislikes: []string{"dislike1"},
	},
}
