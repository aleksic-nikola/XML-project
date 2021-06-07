package handlers

import (
	"log"
	"net/http"
	"xml/profile-serivce/data"
)

type Agents struct {
	l *log.Logger
}

func NewAgents(l *log.Logger) *Agents {
	return &Agents{l}
}

func (u *Agents) GetAgents(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle GET Request for Profiles")

	lp := data.GetAgents()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}