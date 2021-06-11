package service

import (
	"xml/monolit-service/data"
	"xml/monolit-service/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (service *ProductService) CreateProduct(product *data.Product) error {
	error := service.Repo.CreateProduct(product)
	return error
}