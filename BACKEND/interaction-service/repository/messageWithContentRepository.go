package repository

import (
	"fmt"
	"xml/interaction-service/data"

	"gorm.io/gorm"
)

type MessageWithContentRepository struct {
	Database *gorm.DB
}

func (repo *MessageWithContentRepository) CreateMessageWithConent(msgWithContent *data.MessageWithContent) error {
	result := repo.Database.Create(msgWithContent)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}
