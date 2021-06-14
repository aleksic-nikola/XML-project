package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/content-service/data"
	"xml/content-service/service"
)

type StoryHandler struct {
	L       *log.Logger
	Service *service.StoryService
}

func NewStories(l *log.Logger, service *service.StoryService) *StoryHandler {
	return &StoryHandler{l, service}
}

func (handler *StoryHandler) CreateStory(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var story data.Story
	err := story.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(story)

	err = handler.Service.CreateStory(&story)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (p *StoryHandler) GetStories(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request for Posts")

	ls := data.GetStories()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal stories json", http.StatusInternalServerError)
	}
}
