package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/request-service/data"
)

type VerificationRequestRepository struct {
	Database *gorm.DB
}

func (repo *VerificationRequestRepository) CreateVerificationRequest(verificationRequest *data.VerificationRequest) error {
	result := repo.Database.Create(verificationRequest)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}
