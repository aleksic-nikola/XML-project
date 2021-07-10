package service

import (
	"xml/content-service/data"
	"xml/content-service/repository"
)

type CommentService struct {
	Repo *repository.CommentRepository
}

func (service *CommentService) CreateComment(comment *data.Comment) error {
	error := service.Repo.CreateComment(comment)
	return error
}

func (service *CommentService) CommentExists(id uint) (bool, error) {

	exists := service.Repo.CommentExists(id)
	return exists, nil
}
