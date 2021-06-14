package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/content-service/data"
	"xml/content-service/service"
)

type MediaHandler struct {
	L       *log.Logger
	Service *service.MediaService
}

func NewMedias(l *log.Logger, service *service.MediaService) *MediaHandler {
	return &MediaHandler{l, service}
}

func (handler *MediaHandler) CreateMedia(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var media data.Media
	err := media.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(media)

	err = handler.Service.CreateMedia(&media)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}
