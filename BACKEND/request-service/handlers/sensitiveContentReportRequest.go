package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/request-service/data"
	"xml/request-service/service"
)

type SensitiveContentReportRequestHandler struct {
	L *log.Logger
	Service *service.SensitiveContentReportRequestService

}

func NewSensitiveContentReportRequest(l *log.Logger, service *service.SensitiveContentReportRequestService) *SensitiveContentReportRequestHandler {
	return &SensitiveContentReportRequestHandler{l, service}
}

func (handler *SensitiveContentReportRequestHandler) CreateSensitiveContentReportRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var sensitiveContentReportRequest data.SensitiveContentReportRequest
	err := sensitiveContentReportRequest.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(sensitiveContentReportRequest)

	err = handler.Service.CreateSensitiveContentReportRequest(&sensitiveContentReportRequest)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}


func (p *SensitiveContentReportRequestHandler) GetSensitiveContentReportRequests(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request")

	lp := data.GetSensitiveContentReportRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}