package repository

import (
	"fmt"
	"xml/content-service/data"

	"gorm.io/gorm"
)

type PostRepository struct {
	Database *gorm.DB
}

func (repo *PostRepository) CreatePost(post *data.Post) error {
	result := repo.Database.Create(post)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *PostRepository) PostExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Post{}).Count(&count)
	return count != 0
}
