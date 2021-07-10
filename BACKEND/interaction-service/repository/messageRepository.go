package repository

import (
	"fmt"
	"xml/interaction-service/data"

	"gorm.io/gorm"
)

type MessageRepository struct {
	Database *gorm.DB
}

func (repo *MessageRepository) CreateMessage(message *data.Message) error {
	result := repo.Database.Create(message)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}
