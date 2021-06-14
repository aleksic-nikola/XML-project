package service

import (
	"xml/request-service/data"
	"xml/request-service/repository"
)

type FollowRequestService struct {
	Repo *repository.FollowRequestRepository
}

func (service *FollowRequestService) CreateFollowRequest(followRequest *data.FollowRequest) error {
	error := service.Repo.CreateFollowRequest(followRequest)
	return error
}


func (service *FollowRequestService) GetAllRequests() ([]data.FollowRequest, error){
	followReqs, error := service.Repo.GetAllRequests()
	return followReqs, error
}