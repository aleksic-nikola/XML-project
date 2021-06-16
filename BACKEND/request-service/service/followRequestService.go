package service

import (
	"fmt"
	"xml/request-service/data"
	"xml/request-service/repository"
)

type FollowRequestService struct {
	Repo *repository.FollowRequestRepository
}

func (service *FollowRequestService) CreateFollowRequest(followRequest *data.FollowRequest) error {

	 exist, err := service.FollowRequestExists(followRequest.Request.SentBy,followRequest.ForWho)
	 if err != nil{
	 	return fmt.Errorf("Error sql!")
	 }

	 if exist== true{
	 	return fmt.Errorf("Duplicate!")
	 }

	error := service.Repo.CreateFollowRequest(followRequest)
	return error
}


func (service *FollowRequestService) GetAllRequests() ([]data.FollowRequest, error){
	followReqs, error := service.Repo.GetAllRequests()
	return followReqs, error
}

func (service *FollowRequestService) FollowRequestExists(sent_by string, forU string) (bool, error) {

	exists := service.Repo.FollowRequestExists(sent_by, forU)
	return exists, nil
}