package handlers

import (
	"log"
	"net/http"
	"xml/request-service/data"
)

type FollowRequests struct {
	l *log.Logger
}

func NewFollowRequest(l *log.Logger) *FollowRequests {
	return &FollowRequests{l}
}

func (p *FollowRequests) GetFollowRequests(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request")

	lp := data.GetFollowRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
