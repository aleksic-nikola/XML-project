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
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *VerificationRequestRepository) GetInProgressVerificationRequests() data.VerificationRequests {
	var requests data.VerificationRequests
	repo.Database.Where("status = ?", 0).Find(&requests)
	return requests
}

func (repo *VerificationRequestRepository) FindById(id uint) (error, *data.VerificationRequest) {

	var request data.VerificationRequest

	result := repo.Database.Where("id = ? and status = 0", id).First(&request)
	fmt.Println(request)
	if result.RowsAffected != 1 {
 		err := fmt.Errorf("We didnt find any object with that id!");
		return err, &request
	}

	return nil, &request
}

func (repo *VerificationRequestRepository) UpdateVerificationRequest(request *data.VerificationRequest) *data.VerificationRequest {
	repo.Database.Save(&request)
	return request
}

func (repo *VerificationRequestRepository) CheckIfUserHasActiveVR(username string) int64 {
	var request data.VerificationRequest

	result := repo.Database.Where("sent_by = ? and status in (0,1)", username).First(&request)
	fmt.Println("ZAHTEV NADJEN U BAZI:")
	fmt.Println(request)
	fmt.Println("PRONADJENO:")
	fmt.Println(result.RowsAffected)

	return result.RowsAffected
}

