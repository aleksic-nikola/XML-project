package service

import (
	"xml/campaign-service/data"
	"xml/campaign-service/repository"
)

type MultiCampaignService struct {
	Repo *repository.MultiCampaignRepository
}

func (service *MultiCampaignService) CreateMultiCampaign(multiCampaign *data.MultiCampaign) error {
	error := service.Repo.CreateMultiCampaign(multiCampaign)
	return error
}
