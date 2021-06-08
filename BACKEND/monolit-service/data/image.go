package data

import (
	"encoding/json"
	"io"
)

type Image struct {
	Path string `json:"path"`
}

func (p *Image) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Images) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// declaring the collection
type Images []*Image

