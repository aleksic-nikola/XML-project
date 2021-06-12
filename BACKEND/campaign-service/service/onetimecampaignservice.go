package service

import (
	"xml/campaign-service/data"
	"xml/campaign-service/repository"
)

type OneTimeCampaignService struct {
	Repo *repository.OneTimeCampaignRepository
}

func (service *OneTimeCampaignService) CreateOneTimeCampaign(oneTimeCampaign *data.OneTimeCampaign) error {
	error := service.Repo.CreateOneTimeCampaign(oneTimeCampaign)
	return error
}