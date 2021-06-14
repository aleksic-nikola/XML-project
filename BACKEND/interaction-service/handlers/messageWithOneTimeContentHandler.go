package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/interaction-service/data"
	"xml/interaction-service/service"
)

type MessageWithOneTimeContentHandler struct {
	L       *log.Logger
	Service *service.MessageWithOneTimeContentService
}

func NewMessageWithOneTimeContent(l *log.Logger, service *service.MessageWithOneTimeContentService) *MessageWithOneTimeContentHandler {
	return &MessageWithOneTimeContentHandler{l, service}
}

func (handler *MessageWithOneTimeContentHandler) CreateMessageWithOneTimeContent(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var message data.MessageWithOneTimeContent
	err := message.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(message)

	err = handler.Service.CreateMessageWithOneTimeContent(&message)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (m *MessageWithOneTimeContentHandler) GetMessagesWithOneTimeContent(rw http.ResponseWriter, r *http.Request) {
	m.L.Println("Handle GET Request in MessageWithOneTimeContentHandler")

	lp := data.GetMessagesWithOneTimeContent()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
