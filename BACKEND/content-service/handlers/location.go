package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/content-service/data"
	"xml/content-service/service"
)

type LocationHandler struct {
	L       *log.Logger
	Service *service.LocationService
}

func NewLocations(l *log.Logger, service *service.LocationService) *LocationHandler {
	return &LocationHandler{l, service}
}

func (handler *LocationHandler) CreateLocation(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var location data.Location
	err := location.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(location)

	err = handler.Service.CreateLocation(&location)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}
