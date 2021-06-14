package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/request-service/data"
)

type AgentRegistrationRequestRepository struct {
	Database *gorm.DB
}

func (repo *AgentRegistrationRequestRepository) CreateAgentRegistrationRequest(agentRegistrationRequest *data.AgentRegistrationRequest) error {
	result := repo.Database.Create(agentRegistrationRequest)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}


