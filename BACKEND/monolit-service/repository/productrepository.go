package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/monolit-service/data"
)

type ProductRepository struct {
	Database *gorm.DB
}

func (repo *ProductRepository) CreateProduct(product *data.Product) error {
	result := repo.Database.Create(product)
	fmt.Println(result.RowsAffected)
	return result.Error
}
