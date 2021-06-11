package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/search-service/data"
)

type QueryRepository struct {
	Database *gorm.DB
}

func (repo *QueryRepository) CreateQuery(query *data.Query) error {
	result := repo.Database.Create(query)
	fmt.Println(result.RowsAffected)
	return result.Error
}
