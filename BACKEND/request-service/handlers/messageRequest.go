package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/request-service/data"
	"xml/request-service/service"
)

type MessageRequestHandler struct {
	L *log.Logger
	Service *service.MessageRequestService

}

func NewMessageRequest(l *log.Logger, service *service.MessageRequestService) *MessageRequestHandler {
	return &MessageRequestHandler{l, service}
}


func (handler *MessageRequestHandler) CreateMessageRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var messageRequest data.MessageRequest
	err := messageRequest.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(messageRequest)

	err = handler.Service.CreateMessageRequest(&messageRequest)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}


func (p *MessageRequestHandler) GetMessageRequests(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request")

	lp := data.GetMessageRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}