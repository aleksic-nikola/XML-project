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

func (repo *PostRepository) GetPostsByUser(user string) data.Posts {
	var posts data.Posts
	repo.Database.Preload("Likes").Preload("Dislikes").Preload("Medias").Preload("Comments").Where("posted_by = ?", user).Find(&posts)
	return posts
}

func (repo *PostRepository) GetAllPostsForUser(username string) (data.Posts) {
	var posts data.Posts
	repo.Database.Where("posted_by = ?", username).Find(&posts)
	return posts
}

func (repo *PostRepository) GetAllPosts() data.Posts {
	var posts data.Posts
	repo.Database.Preload("Likes").Preload("Dislikes").Preload("Medias").Preload("Comments").Find(&posts)
	return posts
}
