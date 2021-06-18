package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/content-service/data"
	"xml/content-service/data/dtos"
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

func (s *StoryHandler) GetStories(rw http.ResponseWriter, r *http.Request) {
	s.L.Println("Handle GET Request for Posts")

	ls := data.GetStories()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal stories json", http.StatusInternalServerError)
	}
}

func (s *StoryHandler) GetStoriesForCurrentUser(rw http.ResponseWriter, r *http.Request) {
	// send whoami to auth service
	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		s.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var dto dtos.UsernameRole
	err = dto.FromJSON(resp.Body)
	if err != nil {
		
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}
	stories := s.Service.GetAllStoriesForUser(dto.Username)
	err = stories.ToJSON(rw)
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,	
		)
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	//resp, err := http.Get(os.Getenv("profile") + "/whoami")
}
