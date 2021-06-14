package data

import (
	"encoding/json"
	"io"
)

type AgeGroup struct {
	FromAge int `json:"fromage"`
	TillAge int `json:"tillage"`
}

func (ag *AgeGroup) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ag)
}

func (ag *AgeGroups) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ag)
}

type AgeGroups []*AgeGroup