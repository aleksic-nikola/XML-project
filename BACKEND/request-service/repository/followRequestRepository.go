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

func (repo *FollowRequestRepository) GetMyFollowRequests(username string) ([]data.FollowRequest, error) {
	var followReqs []data.FollowRequest
	repo.Database.Where("for_who = ?", username).Find(&followReqs)

	fmt.Println(followReqs)

	return followReqs, nil
}

func (repo *FollowRequestRepository) AcceptFollowRequest(sent_by string, forWho string) error {
	var req data.FollowRequest
	result := repo.Database.Where("sent_by = ? AND for_who = ? AND status=0", sent_by, forWho).Find(&req)

	if(result.RowsAffected!=1){
		return fmt.Errorf("Cant find follow request!")
	}

	fmt.Println("IZ BAZE NASAO REQ: ")
	fmt.Println(req)

	req.Request.Status = data.ACCEPTED

	repo.Database.Save(req)

	return nil
}



