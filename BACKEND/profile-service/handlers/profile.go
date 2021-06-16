package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"xml/profile-service/data"
	dtos2 "xml/profile-service/dto"
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


func (u *ProfileHandler) FollowAccount(rw http.ResponseWriter, r *http.Request){
	//prvo proveravamo ko smo to mi
	jwtToken := r.Header.Get("Authorization")
	resp, err := UserCheck(jwtToken)
	if err != nil {
		u.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	reqBody, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	reqBodyString := string(reqBody)
	fmt.Println("DOBILI: ", reqBodyString)



	bodyBytes, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Fatal(err1)
	}
	bodyString := string(bodyBytes)
	fmt.Println("***********************")
	fmt.Println(bodyString)

	var meDto dtos2.UsernameRoleDto

	json.Unmarshal([]byte(bodyString), &meDto)

	fmt.Println("Ulogovan je: ", meDto.Username)

	//znamo ko smo mi, sada treba da pronadjemo koga treba da zapratimo

	var profileToFollow dtos2.ProfileForFollow

	json.Unmarshal([]byte(reqBodyString), &profileToFollow)

	fmt.Println("Treba da zapratimo: ",profileToFollow.FollowToUsername)

	//sada treba da izvucemo podatke o profilu preko username

	myProfile, errNotFound := u.Service.GetProfileByUsername(meDto.Username)

	if errNotFound!=nil{
		fmt.Println("Not Found: ", meDto.Username)
	}

	profileForFollow, errNotFound := u.Service.GetProfileByUsername(profileToFollow.FollowToUsername)
	if errNotFound!=nil{
		fmt.Println("Not Found: ", meDto.Username)
	}

	fmt.Println("Profil koje treba da se zaprati je JAVAN: ")
	fmt.Println(profileForFollow.PrivacySetting.IsPublic)
	fmt.Println(profileForFollow.Username)

	if(profileForFollow.PrivacySetting.IsPublic == true){
		err = u.Service.FollowProfile(myProfile, profileForFollow)

		if err!=nil{
			http.Error(rw, "Error at Following profile", http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
		return
	}

	err = SendFollowRequest(myProfile, profileForFollow)

	if err!=nil{
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Profil koji sam zapratio prati: ")
	fmt.Println("******************************************************")
	fmt.Println(u.GetAllFollowingByUsername(myProfile.Username))

	myNewProfile, _ := u.Service.GetProfileByUsername("nikola1913")
	fmt.Println("*************** JA PRATIM *****************")
	fmt.Println(myNewProfile.Following)


}

func SendFollowRequest(myProfile *data.Profile, profileToFollow *data.Profile) error {
	client := &http.Client{}


	var dto1 dtos2.FollowRequestDto

	dto1.ForWho = profileToFollow.Username
	dto1.Request.SentBy = myProfile.Username


	json, err := json.Marshal(dto1)

	if err!=nil{
		return fmt.Errorf("Error unmarshaling request json")
	}
	url := "http://localhost:9211/followReqs/add"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

	if err != nil{
		return fmt.Errorf("Error sending to followService req")
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending to request req")
	}
	if res.StatusCode != http.StatusCreated{
		return fmt.Errorf("Duplicate!")
	}



	return nil
}


func (u *ProfileHandler) GetAllFollowingByUsername(username string) []string{
	followingProfiles :=  u.Service.GetAllFollowingByUsername(username)
	fmt.Println(followingProfiles)

	return followingProfiles
}

func (u *ProfileHandler) GetIdByUsername(username string) uint {
	userId, err := u.Service.GetIdByUsername(username)

	if err!=nil{
		fmt.Errorf("Greska.....")
		return 9999
	}

	return userId
}



func UserCheck(tokenString string) (*http.Response, error) {
	err := godotenv.Load()
	if err!=nil{
		fmt.Println("Error at loading env vars\n")
		return nil,err
	}

	client := &http.Client{}
	url := "http://localhost:9090/whoami"
	req, errReq := http.NewRequest("GET", url, nil)

	if errReq != nil{
		return nil, fmt.Errorf("Error with the whoami request")
	}
	req.Header.Add("Authorization", tokenString)
	return client.Do(req)

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