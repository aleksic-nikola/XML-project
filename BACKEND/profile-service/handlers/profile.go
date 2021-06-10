package handlers

import (
	"log"
	"net/http"
	"xml/profile-service/data"
	"xml/profile-service/service"
)

type ProfileHandler struct {
	L *log.Logger
	Service *service.ProfileService
	
}

func NewProfiles(l *log.Logger, service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{l, service}
}

func (u *ProfileHandler) GetProfiles(rw http.ResponseWriter, r *http.Request) {
	u.L.Println("Handle GET Request for Profiles")

	lp := data.GetProfiles()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}
