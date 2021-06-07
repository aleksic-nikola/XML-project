package data

import (
	"encoding/json"
	"io"
	"time"
)

type OneTimeCampaign struct {
	Campaign `json:"campaign"`
	Timestamp time.Time `json:"timestamp"`
}

func (otc *OneTimeCampaign) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(otc)
}

func (otc *OneTimeCampaigns) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(otc)
}

type OneTimeCampaigns []*OneTimeCampaign

func GetOneTimeCampaigns() OneTimeCampaigns {
	return oneTimeCampaignList
}

var oneTimeCampaignList = []*OneTimeCampaign {
	{
		Campaign:
		Campaign{
			ID: 1,
			CreatedBy: "lucyxz",
			Influencers:
			[]string{"a", "b", "c"},
			Ads:
			[]*Ad{
				{
					ID: 1,
					Description: "description",
					Link: "link",
					Product: Product{
						ID:1,
						Availability: 10,
						Name: "Some name",
					},
					Media: Media{
						ID: 1,
						Type: image,
						Path: "path_to_image",
					},
				},
			},
			TargetAudience: TargetAudience{
				ID: 1,
				Tags: []string{"tag1","tag2"},
				AgeGroups: []*AgeGroup {
					{
						FromAge: 15,
						TillAge: 25,
					},
				},
			},
		},
		Timestamp: time.Now(),
	},

}