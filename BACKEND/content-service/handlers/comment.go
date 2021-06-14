package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/content-service/data"
	"xml/content-service/service"
)

type CommentHandler struct {
	L       *log.Logger
	Service *service.CommentService
}

func NewComments(l *log.Logger, service *service.CommentService) *CommentHandler {
	return &CommentHandler{l, service}
}

func (handler *CommentHandler) CreateComment(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var comment data.Comment
	err := comment.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(comment)

	err = handler.Service.CreateComment(&comment)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}
