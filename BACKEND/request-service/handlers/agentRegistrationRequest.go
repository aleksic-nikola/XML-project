package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/request-service/data"
	"xml/request-service/service"
)

type AgentRegistrationRequestHandler struct {
	L*log.Logger
	Service *service.AgentRegistrationRequestService
}

func NewAgentRegistrationRequest(l *log.Logger, service *service.AgentRegistrationRequestService) *AgentRegistrationRequestHandler {
	return &AgentRegistrationRequestHandler{l, service}
}

func (handler *AgentRegistrationRequestHandler) CreateAgentRegistrationRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var agentRegistrationRequest data.AgentRegistrationRequest
	err := agentRegistrationRequest.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(agentRegistrationRequest)

	err = handler.Service.CreateAgentRegistrationRequest(&agentRegistrationRequest)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}





func (p *AgentRegistrationRequestHandler) GetAgentRegistrationRequests(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request")

	lp := data.GetAgentRegistrationRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}