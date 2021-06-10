package handlers

import (
	"log"
	"net/http"
	"xml/profile-service/data"
	"xml/profile-service/service"
)

type VerifiedHandler struct {
	L *log.Logger
	Service *service.VerifiedService
}

func NewVerifieds(l *log.Logger, service *service.VerifiedService) *VerifiedHandler {
	return &VerifiedHandler{l, service}
}

func (u *VerifiedHandler) GetVerifieds(rw http.ResponseWriter, r *http.Request) {
	u.L.Println("Handle GET Request for Profiles")

	lp := data.GetVerifieds()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}