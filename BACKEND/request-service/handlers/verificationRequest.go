package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/request-service/data"
	"xml/request-service/service"
)

type VerificationRequestHandler struct {
	L *log.Logger
	Service *service.VerificationRequestService

}

func NewVerificationRequest(l *log.Logger, service *service.VerificationRequestService) *VerificationRequestHandler {
	return &VerificationRequestHandler{l, service}
}


func (handler *VerificationRequestHandler) CreateVerificationRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var verificationRequest data.VerificationRequest
	err := verificationRequest.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(verificationRequest)

	err = handler.Service.CreateVerificationRequest(&verificationRequest)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}


func (p *VerificationRequestHandler) GetVerificationRequests(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request")

	lp := data.GetVerificationRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

