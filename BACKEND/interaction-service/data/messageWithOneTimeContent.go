package data

import (
	"encoding/json"
	"io"
	"time"
)

type MessageWithOneTimeContent struct {
	MessageID uint    `json:"message_id"`
	Message   Message `json:"message" gorm:"foreignkey=MessageID"`
	Media     Media   `gorm:"embedded"`
}

func (m *MessageWithOneTimeContent) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(e)
}

// collection
type MessagesWithOneTimeContent []*MessageWithOneTimeContent

func (m *MessagesWithOneTimeContent) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func GetMessagesWithOneTimeContent() MessagesWithOneTimeContent {
	return messagesWithOneTimeContentList
}

var messagesWithOneTimeContentList = []*MessageWithOneTimeContent{
	{
		Message: Message{

			From:      "nikola",
			For:       "mark",
			Text:      "pozdrav",
			Timestamp: time.Now(),
		},
		Media: Media{
			ID:   1,
			Path: "mypaaaaaaath",
			Type: image,
		},
	},
}
