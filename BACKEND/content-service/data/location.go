package data

import (
	"encoding/json"
	"io"
)

type Location struct {
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
