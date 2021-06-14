package data

import (
	"encoding/json"
	"io"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID    uint      `json:"post_refer"`
	PostedBy  string    `json:"postedby"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func (c *Comment) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

func (c *Comments) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

type Comments []*Comment

func GetComments() Comments {
	return commentList
}

var commentList = []*Comment{

	{
		//CommentID: 1,
		PostedBy:  "lucyxz",
		Text:      "some text here",
		Timestamp: time.Now(),
	},
}
