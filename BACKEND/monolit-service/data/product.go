package data

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
)

type Product struct {
	gorm.Model
	Name           string        `json:"name"`
	Price          int           `json:"price"`
	Availability   int           `json:"availability"`
	Payments       []Payment     `json:"payment_options" gorm:"many2many:product_payment_options"`
	Images		   []Image		 `json:"images"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// declaring the collection
type Products []*Product

func GetProducts() Products {
	return productList
}

var productList = []*Product{

}

