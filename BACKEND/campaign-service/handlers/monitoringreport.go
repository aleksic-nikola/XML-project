package handlers

import (
	"log"
	"net/http"
	"xml/campaign-service/data"
)

type MonitoringReports struct {
	l *log.Logger
}

func NewMonitoringReports(l *log.Logger) *MonitoringReports {
	return &MonitoringReports{l}
}

func (p *MonitoringReports) GetMonitoringReports(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request for Monitoring Reports")

	ls := data.GetMonitoringReports()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal monitoring reports json" , http.StatusInternalServerError)
	}
}