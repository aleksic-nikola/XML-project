package handlers

import (
	"log"
	"net/http"
	"xml/interaction-service/data"
)

type MessagesWithOneTimeContent struct {
	l *log.Logger
}

func NewMessagesWithOneTimeContent(l *log.Logger) *MessagesWithOneTimeContent {
	return &MessagesWithOneTimeContent{l}
}

func(m *MessagesWithOneTimeContent) GetMessagesWithOneTimeContent(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET Request in MessageWithOneTimeContentHandler")

	lp := data.GetMessagesWithOneTimeContent()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}