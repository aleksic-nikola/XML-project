package data

import (
	"encoding/json"
	"io"
	"time"
)

type MultiCampaign struct {
	CampaignID	uint	  `json:"campaign_id"`
	Campaign    Campaign  `json:"campaign" gorm:"foreignkey:CampaignID"`
	FromDate    time.Time `json:"fromdate" gorm:"type:date"`
	ToDate      time.Time `json:"todate" gorm:"type:date"`
	TimesPerDay int       `json:"timesperday"`
}

func (mc *MultiCampaign) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(mc)
}

func (mc *MultiCampaigns) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(mc)
}

type MultiCampaigns []*MultiCampaign

func GetMultiCampaigns() MultiCampaigns {
	return multiCampaignsList
}

var multiCampaignsList = []*MultiCampaign {
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
					AgeGroup: AgeGroup {
						FromAge: 15,
						TillAge: 25,
					},
				},
			},
			FromDate: time.Now(),
			ToDate: time.Now(),
			TimesPerDay: 3,
	},

}