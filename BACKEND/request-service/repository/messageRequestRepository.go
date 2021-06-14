package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/request-service/data"
)

type MessageRequestRepository struct {
	Database *gorm.DB
}

func (repo *MessageRequestRepository) CreateMessageRequest(messageRequest *data.MessageRequest) error {
	result := repo.Database.Create(messageRequest)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}