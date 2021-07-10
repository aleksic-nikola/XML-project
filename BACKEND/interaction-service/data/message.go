package data

import (
	"encoding/json"
	"io"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	From      string    `json:"from"`
	For       string    `json:"for"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func (m *Message) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(e)
}

// collection
type Messages []*Message

func (m *Messages) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func GetMessages() Messages {
	return messageList
}

var messageList = []*Message{
	{
		From:      "nikola",
		For:       "mark",
		Text:      "pozdrav",
		Timestamp: time.Now(),
	},
	{
		From:      "nikola",
		For:       "danilo",
		Text:      "happy birthday",
		Timestamp: time.Now().AddDate(0, 0, -1),
	},
}
