package handlers

import (
	"log"
	"net/http"
	"xml/profile-serivce/data"
)

type Verifieds struct {
	l *log.Logger
}

func NewVerifieds(l *log.Logger) *Verifieds {
	return &Verifieds{l}
}

func (u *Verifieds) GetVerifieds(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle GET Request for Profiles")

	lp := data.GetVerifieds()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}