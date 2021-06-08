package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID int `json:"id"`
	Availability int `json:"availability"`
	Name string `json:"name"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

type Products []*Product