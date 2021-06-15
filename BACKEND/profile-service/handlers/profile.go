package handlers

import (
	"bytes"
	"encoding/json"
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

	//var userdto dto.UserEditDTO
	var profile dto.ProfileEditDTO

	err := profile.ProfFromJSON(r.Body)

	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(profile)

	oldUsername := "tomlaz"
	err = handler.Service.EditProfileData(profile, oldUsername)

	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}

	// update also user

	requestBody, err := json.Marshal(map[string]string{
		"OldUsername" : oldUsername,
		"Username" : profile.Username,
		"Email" : profile.Email,
		"Name" : profile.Name,
		"LastName" : profile.LastName,
	})

	client := &http.Client{}

	url := "http://localhost:9090/edituser"

	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Errorf("Error with creating new profile")
	}
	_, err = client.Do(req)
	if err != nil {
		fmt.Errorf("Error while  updating user")
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
}


func (handler *ProfileHandler) CreateProfile(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating profile")
	var profile data.Profile

	err := profile.FromJSON(r.Body)

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

func (handler *ProfileHandler) EditProfilePrivacySettings(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("updating privacy settings")
	var privacySettings data.PrivacySetting

	err := privacySettings.FromJSON(r.Body)

	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(privacySettings)

	// get username from current session
	username := "tomlaz"
	err = handler.Service.EditProfilePrivacySettings(privacySettings, username)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
}

func (handler *ProfileHandler) EditProfileNotificationSettings(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("updating notification settings")
	var notifSettings data.NotificationSetting

	err := notifSettings.FromJSON(r.Body)

	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(notifSettings)

	// get username from current session
	username := "tomlaz"
	err = handler.Service.EditProfileNotificationSettings(notifSettings, username)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
}



