package repository

import (
	"fmt"
	"xml/profile-service/data"

	"gorm.io/gorm"
)
type AgentRepository struct {
	Database *gorm.DB
}

func (repo *AgentRepository) CreateAgent(user *data.Agent) error {
	result := repo.Database.Create(user)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *AgentRepository) AgentExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Profile{}).Count(&count)
	return count != 0
}