package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/campaign-service/data"
	"xml/campaign-service/service"
)

type OneTimeCampaignHandler struct {
	L *log.Logger
	Service *service.OneTimeCampaignService
}

func NewOneTimeCampaigns(l *log.Logger, service *service.OneTimeCampaignService) *OneTimeCampaignHandler {
	return &OneTimeCampaignHandler{l, service}
}

func (handler *OneTimeCampaignHandler) CreateOneTimeCampaign(rw http.ResponseWriter, r *http.Request)  {
	fmt.Println("creating one time campaign")
	var oneTimeCampaign data.OneTimeCampaign
	err := oneTimeCampaign.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(oneTimeCampaign)

	err = handler.Service.CreateOneTimeCampaign(&oneTimeCampaign)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (otc *OneTimeCampaignHandler) GetOneTimeCampaigns(rw http.ResponseWriter, r *http.Request) {
	otc.L.Println("Handle GET Request for One Time Campaigns")

	ls := data.GetOneTimeCampaigns()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal monitoring reports json" , http.StatusInternalServerError)
	}
}
