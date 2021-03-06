package data

import (
	"encoding/json"
	"io"
)

type Media struct {
	Type MediaType `json:"type" gorm:"type:text"`
	Path string    `json:"path"`
}

func (m *Media) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(m)
}

func (m *Medias) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

type Medias []*Media
