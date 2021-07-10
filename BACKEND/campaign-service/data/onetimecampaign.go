package data

import (
	"encoding/json"
	"io"
	"time"
)

type OneTimeCampaign struct {
	CampaignID uint 	 `json:"campaign_id"`
	Campaign   Campaign  `json:"campaign" gorm:"foreignkey:CampaignID"`
	Timestamp  time.Time `json:"timestamp" gorm:"type:date"`
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
			CreatedBy: "lucyxz",
			Influencers:
			[]string{"a", "b", "c"},
			Ads: []Ad{
				{
					ID: 1,
					Description: "description",
					Link: "link",
					Product: Product{
						Availability: 10,
						Name: "Some name",
					},
					Media: Media{
						Type: image,
						Path: "path_to_image",
					},
				},
			},
			TargetAudience: TargetAudience{
				Tags: []string{"tag1","tag2"},
				AgeGroup: AgeGroup {
					FromAge: 15,
					TillAge: 25,
				},
			},
		},
		Timestamp: time.Now(),
	},

}