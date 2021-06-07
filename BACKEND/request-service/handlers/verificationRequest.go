package handlers

import (
	"log"
	"net/http"
	"xml/request-service/data"
)

type VerificationRequests struct {
	l *log.Logger
}

func NewVerificationRequest(l *log.Logger) *VerificationRequests {
	return &VerificationRequests{l}
}

func (p *VerificationRequests) GetVerificationRequests(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request")

	lp := data.GetVerificationRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}