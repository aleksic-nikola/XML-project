package data

import (
	"encoding/json"
	"io"
)

type TargetAudience struct {
	Tags     []string `json:"tags" gorm:"type:text"`
	AgeGroup AgeGroup `json:"age_group" gorm:"embedded"`
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