package data

import (
	"encoding/json"
	"io"
)

type InfluenceRequest struct {

	RequestID uint `json:"request_id"`
	Request Request `json:"request" gorm:"foreignkey=RequestID"`
	CampaignID string `json:"campaignid"`

}


func (p *InfluenceRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// collection
type InfluenceRequests []*InfluenceRequest

// encode (using json new encoder over marshall)
func (p *InfluenceRequests) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetInfluenceRequests() InfluenceRequests {
	return influenceRequestList
}


var influenceRequestList = []*InfluenceRequest{

	{
		RequestID: 4,
		Request : Request{
			//ID: 2,
			SentBy : "wintzy",
			Status: ACCEPTED,
		},
		CampaignID: "34123",

	},
	{
		RequestID: 5,
		Request : Request{
			//ID: 2,
			SentBy : "dani",
			Status: INPROCESS,
		},
		CampaignID: "12123",
	},
}