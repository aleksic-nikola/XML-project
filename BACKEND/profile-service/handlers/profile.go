package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"xml/profile-service/data"
	"xml/profile-service/service"
)

type ProfileHandler struct {
	L *log.Logger
	Service *service.ProfileService
	
}

func NewProfiles(l *log.Logger, service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{l, service}
}

func (u *ProfileHandler) GetProfiles(rw http.ResponseWriter, r *http.Request) {
	u.L.Println("Handle GET Request for Profiles")

	lp := data.GetProfiles()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}

func (u *ProfileHandler) IsUserPublic(rw http.ResponseWriter, request *http.Request)  {
	params := mux.Vars(request)
	username := params["username"]

	dto, err := u.Service.IsUserPublic(username)

	if err != nil {
		http.Error(rw, "Unable to find user with that username", http.StatusNotFound)
		return
	}

	err = dto.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal posts json", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) CreateProfile(rw http.ResponseWriter, request *http.Request) {
	fmt.Println("creating user")
	var profile data.Profile
	err := profile.FromJSON(request.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(profile)

	err = handler.Service.CreateProfile(&profile)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}
