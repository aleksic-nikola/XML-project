package data

import (
	"encoding/json"
	"io"
)

type Campaign struct {
	ID int `json:"id"`
	CreatedBy string `json:"createdby"`
	Influencers []string `json:"influencers"`
	Ads `json:"ads"`
	TargetAudience `json:"targetaudiences"`
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