package repository

import (
	"fmt"
	"gorm.io/gorm"
	"xml/request-service/data"

)

type FollowRequestRepository struct {
	Database *gorm.DB
}

func (repo *FollowRequestRepository) CreateFollowRequest(followRequest *data.FollowRequest) error {
	result := repo.Database.Create(followRequest)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

//-------------------------------------------------------------------

func (repo *FollowRequestRepository) GetAllRequests() ([]data.FollowRequest, error){
	var followReqs []data.FollowRequest
	result := repo.Database.Find(&followReqs)
	fmt.Println("REZULTAT MEM: ");
	fmt.Println(followReqs)
	fmt.Println("REZULTAT rowsAffected: ");

	fmt.Println(result.RowsAffected)

	return followReqs, result.Error

}

func (repo *FollowRequestRepository) FollowRequestExists(sent_by string, forU string) bool {
	var count int64
	repo.Database.Where("sent_by = ? AND for_who = ? AND status=0", sent_by, forU).Find(&data.FollowRequest{}).Count(&count)
	return count != 0
}