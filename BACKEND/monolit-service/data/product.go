package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID             int           `json:"id"`
	Price          int           `json:"price"`
	Availability   int           `json:"availability"`
	PaymentOptions []PaymentOption `json:"payment_options"`
	Image		   []Image		 `json:"image"`
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
	{
		ID: 1,
		Price: 300,
		Availability: 50,
		PaymentOptions: []PaymentOption {
			ONPICKUP,
		},
		Image: []Image {
			{
				Path: "putanja",
			},
		},
	},
}

