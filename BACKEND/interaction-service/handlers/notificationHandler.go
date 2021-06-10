package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/interaction-service/data"
	"xml/interaction-service/service"
)

type NotificationHandler struct {
	L       *log.Logger
	Service *service.NotificationService
}

func NewNotifications(l *log.Logger, service *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{l, service}
}

func (handler *NotificationHandler) CreateNotification(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var notification data.Notification
	err := notification.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(notification)

	err = handler.Service.CreateNotification(&notification)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (m *NotificationHandler) GetNotifications(rw http.ResponseWriter, r *http.Request) {
	m.L.Println("Handle GET Request in MessageHandler")

	lp := data.GetMessages()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
