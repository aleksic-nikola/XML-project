package handlers

import (
	"log"
	"net/http"
	"xml/profile-serivce/data"
)

type Profiles struct {
	l *log.Logger
}

func NewProfiles(l *log.Logger) *Profiles {
	return &Profiles{l}
}

func (u *Profiles) GetProfiles(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle GET Request for Profiles")

	lp := data.GetProfiles()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}
