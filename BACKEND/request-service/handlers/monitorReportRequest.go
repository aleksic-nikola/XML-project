package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/request-service/data"
	"xml/request-service/service"
)

type MonitorReportRequestHandler struct {
	L *log.Logger
	Service *service.MonitorReportRequestService

}

func NewMonitorReportRequest(l *log.Logger, service *service.MonitorReportRequestService) *MonitorReportRequestHandler {
	return &MonitorReportRequestHandler{l, service}
}


func (handler *MonitorReportRequestHandler) CreateMonitorReportRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var monitorReportRequest data.MonitorReportRequest
	err := monitorReportRequest.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(monitorReportRequest)

	err = handler.Service.CreateMonitorReportRequest(&monitorReportRequest)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}



func (p *MonitorReportRequestHandler) GetMonitorReportRequests(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request")

	lp := data.GetMonitorReportRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
