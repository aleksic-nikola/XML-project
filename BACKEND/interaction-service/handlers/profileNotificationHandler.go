package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/interaction-service/data"
	"xml/interaction-service/service"
)

type ProfileNotificationHandler struct {
	L       *log.Logger
	Service *service.ProfileNotificationService
}

func NewProfileNotifications(l *log.Logger, service *service.ProfileNotificationService) *ProfileNotificationHandler {
	return &ProfileNotificationHandler{l, service}
}

func (handler *ProfileNotificationHandler) CreateProfileNotification(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var notif data.ProfileNotification
	err := notif.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(notif)

	err = handler.Service.CreateProfileNotification(&notif)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (p *ProfileNotificationHandler) GetProfileNotifications(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request in profileNotificationHandler")

	lp := data.GetPostNotifications()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
