package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
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

	var profile dto.ProfileEditDTO

	// send whoami to auth service
	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		handler.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var dto dto.UsernameRole

	fmt.Println("-->ULOGOVAN SAM KAO: ")
	err = dto.URFromJSON(resp.Body)

	fmt.Println("Logged as: " + dto.Username + " ,with role:" +dto.Role)
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	//var userdto dto.UserEditDTO

	fmt.Println("===========HERE I AM==========")
	err = profile.ProfFromJSON(r.Body)

	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(profile)

	oldUsername := dto.Username
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

	// send whoami to auth service
	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		handler.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var dto dto.UsernameRole

	fmt.Println("-->ULOGOVAN SAM KAO: ")
	err = dto.URFromJSON(resp.Body)

	fmt.Println("Logged as: " + dto.Username + " ,with role:" +dto.Role)
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	err = privacySettings.FromJSON(r.Body)

	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(privacySettings)

	// get username from current session
	username := dto.Username
	err = handler.Service.EditProfilePrivacySettings(privacySettings, username)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
}

func (handler *ProfileHandler) EditProfileNotificationSettings(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("updating notification settings")
	var notifSettings data.NotificationSetting

	// send whoami to auth service
	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		handler.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var dto dto.UsernameRole

	fmt.Println("-->ULOGOVAN SAM KAO: ")
	err = dto.URFromJSON(resp.Body)

	fmt.Println("Logged as: " + dto.Username + " ,with role:" +dto.Role)
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	err = notifSettings.FromJSON(r.Body)

	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(notifSettings)

	// get username from current session
	username := dto.Username
	err = handler.Service.EditProfileNotificationSettings(notifSettings, username)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
}

func UserCheck(tokenString string) (*http.Response, error) {

	godotenv.Load()
	client := &http.Client{}
	url := "http://" + GetVariable("auth") + "/whoami"
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error with the whoami request")
	}
	req.Header.Add("Authorization", tokenString)
	return client.Do(req)
}



