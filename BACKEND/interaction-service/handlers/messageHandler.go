package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/interaction-service/data"
	"xml/interaction-service/service"
)

type MessageHandler struct {
	L       *log.Logger
	Service *service.MessageService
}

func NewMessages(l *log.Logger, service *service.MessageService) *MessageHandler {
	return &MessageHandler{l, service}
}

func (handler *MessageHandler) CreateMessage(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var message data.Message
	err := message.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(message)

	err = handler.Service.CreateMessage(&message)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (m *MessageHandler) GetMessages(rw http.ResponseWriter, r *http.Request) {
	m.L.Println("Handle GET Request in MessageHandler")

	lp := data.GetMessages()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
