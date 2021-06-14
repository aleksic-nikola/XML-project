package data

import (
	"encoding/json"
	"io"
)

type Ad struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Product     Product `json:"products" gorm:"embedded"`
	Media       Media   `json:"media" gorm:"embedded"`
}

func (a *Ad) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func (a *Ads) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

type Ads []*Ad