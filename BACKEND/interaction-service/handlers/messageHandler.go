package handlers

import (
	"log"
	"net/http"
	"xml/interaction-service/data"
)

type Messages struct {
	l *log.Logger
}

func NewMessages(l *log.Logger) *Messages {
	return &Messages{l}
}

func(m *Messages) GetMessages(rw http.ResponseWriter, r *http.Request) {
	m.l.Println("Handle GET Request in MessageHandler")

	lp := data.GetMessages()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}