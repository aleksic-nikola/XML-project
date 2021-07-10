package service

import (
	"xml/content-service/data"
	"xml/content-service/repository"
)

type LocationService struct {
	Repo *repository.LocationRepository
}

func (service *LocationService) CreateLocation(location *data.Location) error {
	error := service.Repo.CreateLocation(location)
	return error
}

func (service *LocationService) LocationExists(id uint) (bool, error) {

	exists := service.Repo.LocationExists(id)
	return exists, nil
}
