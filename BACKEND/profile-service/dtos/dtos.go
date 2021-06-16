package dtos

import (
	"encoding/json"
	"io"
)

type ProfilePublic struct {
	Public bool `json:"ispublic"`
}

func (p *ProfilePublic) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *ProfilePublic) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}