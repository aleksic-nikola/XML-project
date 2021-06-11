package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/monolit-service/data"
	"xml/monolit-service/service"
)


type ProductHandler struct {
	L *log.Logger
	Service *service.ProductService
}

func NewProducts(l *log.Logger, service *service.ProductService) *ProductHandler {
	return &ProductHandler{l, service}
}

func (handler *ProductHandler) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating product")
	var product data.Product
	err := product.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(product)

	err = handler.Service.CreateProduct(&product)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (p *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request for Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}