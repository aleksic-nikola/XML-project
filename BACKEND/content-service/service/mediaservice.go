package service

import (
	"xml/content-service/data"
	"xml/content-service/repository"
)

type MediaService struct {
	Repo *repository.MediaRepository
}

func (service *MediaService) CreateMedia(media *data.Media) error {
	error := service.Repo.CreateMedia(media)
	return error
}

func (service *MediaService) MediaExists(id uint) (bool, error) {

	exists := service.Repo.MediaExists(id)
	return exists, nil
}
