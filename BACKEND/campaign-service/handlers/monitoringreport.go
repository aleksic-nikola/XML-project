package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/campaign-service/data"
	"xml/campaign-service/service"
)

type MonitoringReportHandler struct {
	L *log.Logger
	Service *service.MonitoringReportService
}

func NewMonitoringReports(l *log.Logger, service *service.MonitoringReportService) *MonitoringReportHandler {
	return &MonitoringReportHandler{l, service}
}

func (handler *MonitoringReportHandler) CreateMonitoringReport(rw http.ResponseWriter, r *http.Request)  {
	fmt.Println("creating monitoring report")
	var monitoringReport data.MonitoringReport
	err := monitoringReport.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(monitoringReport)

	err = handler.Service.CreateMonitoringReport(&monitoringReport)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (p *MonitoringReportHandler) GetMonitoringReports(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request for Monitoring Reports")

	ls := data.GetMonitoringReports()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal monitoring reports json" , http.StatusInternalServerError)
	}
}