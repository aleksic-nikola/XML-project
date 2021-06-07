package data

import (
	"encoding/json"
	"io"
)

type TargetAudience struct {
	ID int `json:"id"`
	Tags []string `json:"tags"`
	AgeGroups
}

func (ta *TargetAudience) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ta)
}

func (ta *TargetAudiences) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ta)
}

type TargetAudiences []*TargetAudience