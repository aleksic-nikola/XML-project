package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/request-service/data"
	"xml/request-service/service"
)

type InfluenceRequestHandler struct {
	L *log.Logger
	Service *service.InfluenceRequestService

}

func NewInfluenceRequest(l *log.Logger, service *service.InfluenceRequestService) *InfluenceRequestHandler {
	return &InfluenceRequestHandler{l, service}
}


func (handler *InfluenceRequestHandler) CreateInfluenceRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var influenceRequest data.InfluenceRequest
	err := influenceRequest.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(influenceRequest)

	err = handler.Service.CreateInfluenceRequest(&influenceRequest)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}


func (p *InfluenceRequestHandler) GetInfluenceRequests(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request")


	lp := data.GetInfluenceRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
