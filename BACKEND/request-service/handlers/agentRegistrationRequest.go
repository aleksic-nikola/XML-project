package handlers

import (
	"log"
	"net/http"
	"xml/request-service/data"
)

type AgentRegistrationRequests struct {
	l *log.Logger
}

func NewAgentRegistrationRequest(l *log.Logger) *AgentRegistrationRequests {
	return &AgentRegistrationRequests{l}
}

func (p *AgentRegistrationRequests) GetAgentRegistrationRequests(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Request")

	lp := data.GetAgentRegistrationRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}