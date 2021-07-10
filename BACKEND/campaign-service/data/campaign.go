package data

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
)

type Campaign struct {
	gorm.Model
	CreatedBy      string         `json:"createdby"`
	Influencers    []string       `json:"influencers" gorm:"type:text"`
	Ads            []Ad           `json:"ads" gorm:"many2many:campaign_ads"`
	TargetAudience TargetAudience `json:"targetaudiences" gorm:"embedded"`
}

func (c *Campaign) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

func (c *Campaigns) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

type Campaigns []*Campaign