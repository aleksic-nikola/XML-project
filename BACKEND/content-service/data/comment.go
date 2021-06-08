package data

import (
	"encoding/json"
	"io"
	"time"
)

type Comment struct {
	ID int `json:"id"`
	PostedBy string `json:"postedby"`
	Text string `json:"text"`
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
		ID: 1,
		PostedBy: "lucyxz",
		Text: "some text here",
		Timestamp: time.Now(),
	},
}