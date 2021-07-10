package handlers

import (
	"log"
	"net/http"
	"xml/profile-service/data"
	"xml/profile-service/service"
)

type AgentHandler struct {
	L *log.Logger
	Service *service.AgentService
}

func NewAgents(l *log.Logger, service *service.AgentService) *AgentHandler {
	return &AgentHandler{l, service}
}

func (u *AgentHandler) GetAgents(rw http.ResponseWriter, r *http.Request) {
	u.L.Println("Handle GET Request for Profiles")

	lp := data.GetAgents()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}