package data

import (
	"encoding/json"
	"io"
	"time"
)

type MessageWithContent struct {
	Message Message `json:"message"`
	ContentID int `json:"contentid"`
	//ContentType ContentType `json:"ticketstate"`
}

func(m *MessageWithContent) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(e)
}

// collection
type MessagesWithContent []*MessageWithContent

func(m *MessagesWithContent) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func GetMessagesWithContent() MessagesWithContent {
	return messagesWithContentList
}

var messagesWithContentList = []*MessageWithContent{
	{
		Message: Message{
			ID:        1,
			From:      "nikola",
			For:       "mark",
			Text:      "pozdrav",
			Timestamp: time.Now(),
		},
		ContentID:   1,
	},
	{
		Message: Message{
			ID: 1,
			From: "nikola",
			For: "danilo",
			Text: "happy birthday",
			Timestamp: time.Now().AddDate(0, 0, -1),
		},
		ContentID:   1,
	},

}