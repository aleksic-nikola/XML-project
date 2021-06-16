package service

import (
	"xml/content-service/data"
	"xml/content-service/repository"
)

type PostService struct {
	Repo *repository.PostRepository
}

func (service *PostService) CreatePost(post *data.Post) error {
	error := service.Repo.CreatePost(post)
	return error
}

func (service *PostService) PostExists(id uint) (bool, error) {

	exists := service.Repo.PostExists(id)
	return exists, nil
}

func (service *PostService) GetPostsByUser(username string) data.Posts {
	posts := service.Repo.GetPostsByUser(username)
	return posts
}
