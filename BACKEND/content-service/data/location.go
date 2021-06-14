package data

import (
	"encoding/json"
	"io"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	//LocationID int    `json:"id" gorm:"primaryKey"`
	Country string `json:"country"`
	City    string `json:"city"`
}

func (l *Location) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(l)
}

func (l *Location) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}
