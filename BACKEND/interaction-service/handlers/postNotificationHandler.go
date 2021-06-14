package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/interaction-service/data"
	"xml/interaction-service/service"
)

type PostNotificationHandler struct {
	L       *log.Logger
	Service *service.PostNotificationService
}

func NewPostNotifications(l *log.Logger, service *service.PostNotificationService) *PostNotificationHandler {
	return &PostNotificationHandler{l, service}
}

func (handler *PostNotificationHandler) CreatePostNotification(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var notif data.PostNotification
	err := notif.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(notif)

	err = handler.Service.CreatePostNotification(&notif)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (p *PostNotificationHandler) GetPostNotifications(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request in postNotificationHandler")

	lp := data.GetPostNotifications()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
