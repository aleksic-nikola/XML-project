package repository

import (
	"fmt"
	"xml/interaction-service/data"

	"gorm.io/gorm"
)

type MessageWithOneTimeContentRepository struct {
	Database *gorm.DB
}

func (repo *MessageWithOneTimeContentRepository) CreateMessageWithOneTimeConent(msgWithOneTimeContent *data.MessageWithOneTimeContent) error {
	result := repo.Database.Create(msgWithOneTimeContent)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}
