package handlers

import (
	"log"
	"net/http"
	"xml/request-service/data"
)

type SensitiveContentReportRequests struct {
	l *log.Logger
}

func NewSensitiveContentReportRequest(l *log.Logger) *SensitiveContentReportRequests {
	return &SensitiveContentReportRequests{l}
}

func (p *SensitiveContentReportRequests) GetSensitiveContentReportRequests(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request")

	lp := data.GetSensitiveContentReportRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}