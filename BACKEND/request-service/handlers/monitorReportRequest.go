package handlers

import (
	"log"
	"net/http"
	"xml/request-service/data"
)

type MonitorReportRequests struct {
	l *log.Logger
}

func NewMonitorReportRequest(l *log.Logger) *MonitorReportRequests {
	return &MonitorReportRequests{l}
}

func (p *MonitorReportRequests) GetMonitorReportRequests(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request")

	lp := data.GetMonitorReportRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
