package handlers

import (
	"log"
	"net/http"
	"xml/interaction-service/data"
)

type MessagesWithContent struct {
	l *log.Logger
}

func NewMessagesWithContent(l *log.Logger) *MessagesWithContent {
	return &MessagesWithContent{l}
}

func(m *MessagesWithContent) GetMessagesWithContent(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET Request in MessageWithContentHandler")

	lp := data.GetMessagesWithContent()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}