package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/interaction-service/data"
	"xml/interaction-service/service"
)

type MessageWithContentHandler struct {
	L       *log.Logger
	Service *service.MessageWithContentService
}

func NewMessagesWithContent(l *log.Logger, service *service.MessageWithContentService) *MessageWithContentHandler {
	return &MessageWithContentHandler{l, service}
}

func (handler *MessageWithContentHandler) CreateMessageWithContent(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var message data.MessageWithContent
	err := message.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(message)

	err = handler.Service.CreateMessageWithContent(&message)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (m *MessageWithContentHandler) GetMessagesWithContent(rw http.ResponseWriter, r *http.Request) {
	m.L.Println("Handle GET Request in MessageWithContentHandler")

	lp := data.GetMessagesWithContent()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
