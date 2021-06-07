package handlers

import (
	"log"
	"net/http"
	"xml/campaign-service/data"
)

type OneTimeCampaigns struct {
	l *log.Logger
}

func NewOneTimeCampaigns(l *log.Logger) *OneTimeCampaigns {
	return &OneTimeCampaigns{l}
}

func (otc *OneTimeCampaigns) GetOneTimeCampaigns(rw http.ResponseWriter, r *http.Request) {
	otc.l.Println("Handle GET Request for One Time Campaigns")

	ls := data.GetOneTimeCampaigns()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal monitoring reports json" , http.StatusInternalServerError)
	}
}
