package service

import (
	"xml/search-service/data"
	"xml/search-service/repository"
)

type QueryService struct {
	Repo *repository.QueryRepository
}

func (service *QueryService) CreateQuery(query *data.Query) error {
	error := service.Repo.CreateQuery(query)
	return error
}
