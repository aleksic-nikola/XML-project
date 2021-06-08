package handlers

import (
	"log"
	"net/http"
	"xml/campaign-service/data"
)

type MultiCampaigns struct {
	l *log.Logger
}

func NewMultiCampaigns(l *log.Logger) *MultiCampaigns {
	return &MultiCampaigns{l}
}

func (mc *MultiCampaigns) GetMultiCampaigns(rw http.ResponseWriter, r *http.Request) {
	mc.l.Println("Handle GET Request for One Time Campaigns")

	ls := data.GetMultiCampaigns()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal monitoring reports json" , http.StatusInternalServerError)
	}
}