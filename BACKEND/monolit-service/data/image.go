package data

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
)

type Image struct {
	gorm.Model
	Path string `json:"path"`
	ProductID uint `json:"product_refer"`
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

