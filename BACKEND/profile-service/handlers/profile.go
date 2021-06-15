package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/profile-service/data"
	"xml/profile-service/dto"
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

func (handler *ProfileHandler) EditProfileData(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("updating profile data + user data")

	var userdto dto.UserEditDTO
	var profiledto dto.ProfileEditDTO

	err := profiledto.ProfFromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(profiledto)
	// TODO: with register -> create also a new profile
	//
	// first check updating -- SEND USER UPDATE
	// send Profile update
	// if everything OK -> save them


}
