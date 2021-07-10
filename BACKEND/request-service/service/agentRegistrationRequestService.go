package service

import (
	"xml/request-service/data"
	"xml/request-service/repository"
)

type AgentRegistrationRequestService struct {
	Repo *repository.AgentRegistrationRequestRepository
}


func (service *AgentRegistrationRequestService) CreateAgentRegistrationRequest(agentRegistrationRequest *data.AgentRegistrationRequest) error {
	error := service.Repo.CreateAgentRegistrationRequest(agentRegistrationRequest)
	return error
}
