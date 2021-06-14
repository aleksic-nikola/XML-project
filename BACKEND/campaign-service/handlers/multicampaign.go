package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/campaign-service/data"
	"xml/campaign-service/service"
)

type MultiCampaignHandler struct {
	L *log.Logger
	Service *service.MultiCampaignService
}

func NewMultiCampaigns(l *log.Logger, service *service.MultiCampaignService) *MultiCampaignHandler {
	return &MultiCampaignHandler{l, service}
}

func (handler *MultiCampaignHandler) CreateMultiCampaign(rw http.ResponseWriter, r *http.Request)  {
	fmt.Println("creating multi campaign")
	var multiCampaign data.MultiCampaign
	err := multiCampaign.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(multiCampaign)

	err = handler.Service.CreateMultiCampaign(&multiCampaign)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (mc *MultiCampaignHandler) GetMultiCampaigns(rw http.ResponseWriter, r *http.Request) {
	mc.L.Println("Handle GET Request for Multi Campaigns")

	ls := data.GetMultiCampaigns()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal monitoring reports json" , http.StatusInternalServerError)
	}
}