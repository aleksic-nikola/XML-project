package handlers

import (
	"log"
	"net/http"
	"xml/profile-serivce/data"
)

type Admins struct {
	l *log.Logger
}

func NewAdmins(l *log.Logger) *Admins {
	return &Admins{l}
}

func (u *Admins) GetAdmins(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle GET Request for Profiles")

	lp := data.GetAdmins()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}