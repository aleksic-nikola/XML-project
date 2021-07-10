package repository

import (
	"fmt"
	"xml/content-service/data"

	"gorm.io/gorm"
)

type CommentRepository struct {
	Database *gorm.DB
}

func (repo *CommentRepository) CreateComment(comment *data.Comment) error {
	result := repo.Database.Create(comment)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *CommentRepository) CommentExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Comment{}).Count(&count)
	return count != 0
}
