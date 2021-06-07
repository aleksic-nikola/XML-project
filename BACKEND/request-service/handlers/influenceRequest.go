package handlers

import (
	"log"
	"net/http"
	"xml/request-service/data"
)

type InfluenceRequests struct {
	l *log.Logger
}

func NewInfluenceRequest(l *log.Logger) *InfluenceRequests {
	return &InfluenceRequests{l}
}

func (p *InfluenceRequests) GetInfluenceRequests(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request")

	lp := data.GetInfluenceRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
