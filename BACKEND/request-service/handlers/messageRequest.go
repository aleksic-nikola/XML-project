package handlers

import (
	"log"
	"net/http"
	"xml/request-service/data"
)

type MessageRequests struct {
	l *log.Logger
}

func NewMessageRequest(l *log.Logger) *MessageRequests {
	return &MessageRequests{l}
}

func (p *MessageRequests) GetMessageRequests(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request")

	lp := data.GetMessageRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}