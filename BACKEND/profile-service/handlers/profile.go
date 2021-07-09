package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"xml/profile-service/constants"
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

func (u *ProfileHandler) IsUserPublic(rw http.ResponseWriter, request *http.Request)  {
	params := mux.Vars(request)
	username := params["username"]
	fmt.Println(username)

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

	//url := "http://localhost:9090/edituser"
	url := "http://" + constants.AUTH_SERVICE_URL + "/edituser"

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





func (u *ProfileHandler) FollowAccount(rw http.ResponseWriter, r *http.Request){
	fmt.Println("************** USLI U FOLLOW ******************")

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

	var meDto dto.UsernameRoleDto

	json.Unmarshal([]byte(bodyString), &meDto)

	fmt.Println("Ulogovan je: ", meDto.Username)

	//znamo ko smo mi, sada treba da pronadjemo koga treba da zapratimo

	var profileToFollow dto.ProfileForFollow

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

	fmt.Println("Ja pratim: ")
}



func SendFollowRequest(myProfile *data.Profile, profileToFollow *data.Profile) error {
	client := &http.Client{}


	var dto1 dto.FollowRequestDto

	dto1.ForWho = profileToFollow.Username
	dto1.Request.SentBy = myProfile.Username


	json, err := json.Marshal(dto1)


	if err!=nil{
		return fmt.Errorf("Error unmarshaling request json")
	}
	url := "http://" + constants.REQUEST_SERVICE_URL + "/followReqs/add"

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

func (u *ProfileHandler) GetIdByUsername(username string) uint {
	userId, err := u.Service.GetIdByUsername(username)

	if err!=nil{
		fmt.Errorf("Greska.....")
		return 9999
	}

	return userId
}


func BodyToJson(body io.ReadCloser) (string, error) {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return "", fmt.Errorf("Error with reading body...")
	}
	return string(bodyBytes), nil
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


func (u *ProfileHandler) GetAllFollowingByUsername(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	//jsonString, err := BodyToJson(r.Body)
	rw.Header().Set("Content-Type", "application/json")
	var usernameDto dto.UsernameDto

	fmt.Println("JSON KOJI SMO DOBILI: ")
	fmt.Println(r.Body)

	err := usernameDto.FromJSON(r.Body)

	if err!=nil{
		fmt.Println("Error FromJSON...")
		return
	}

	/*
	err = json.Unmarshal([]byte(jsonString), &usernameDto)
	if err != nil {
		fmt.Println("Error at Unmarshal")
		return
	}
*/
	followingProfiles :=  u.Service.GetAllFollowingByUsername(usernameDto.Username)
	followingProfilesJson, _ := json.Marshal(followingProfiles)

	_, err1 := rw.Write(followingProfilesJson)
	if err1!=nil{
		fmt.Println("Error with Write!")
		return
	}
}

func (u *ProfileHandler) GetAllFollowersByUsername(rw http.ResponseWriter, r *http.Request) {

	//jsonString, err := BodyToJson(r.Body)
	rw.Header().Set("Content-Type", "application/json")
	var usernameDto dto.UsernameDto

	fmt.Println("JSON KOJI SMO DOBILI: ")
	fmt.Println(r.Body)

	err := usernameDto.FromJSON(r.Body)

	if err!=nil{
		fmt.Println("Error FromJSON...")
		return
	}
	//var usernameDto dto.UsernameDto
	//fmt.Println("DOBIOOO: ")
	//fmt.Println(jsonString)
	//err = json.Unmarshal([]byte(jsonString), &usernameDto)
	//if err != nil {
	//	fmt.Println("Error at Unmarshal")
	//	http.Error(rw, "Error at Unmarshal", http.StatusBadRequest)
	//	return
	//}

	followersProfiles :=  u.Service.GetAllFollowersByUsername(usernameDto.Username)
	followersProfilesJson, _ := json.Marshal(followersProfiles)


	_, err1 := rw.Write(followersProfilesJson)
	if err1!=nil{
		fmt.Println("Error with Write!")
		return
	}
}





func (u *ProfileHandler) AcceptFollow(rw http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	resp, err := UserCheck(jwtToken)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	reqBody, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	reqBodyString := string(reqBody)

	fmt.Println("DOBIO OVDE OVAKAV JSON: ")
	fmt.Println(reqBodyString)

	bodyBytes, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Fatal(err1)
	}
	bodyString := string(bodyBytes)

	var meDto dto.UsernameRoleDto

	err = json.Unmarshal([]byte(bodyString), &meDto)

	if err!=nil{
		http.Error(rw, "Can't unmarshal meDto body!", http.StatusBadRequest)
		return
	}

	fmt.Println("Ulogovan je: ", meDto.Username)

	var profileToFollow dto.ProfileForFollow

	err = json.Unmarshal([]byte(reqBodyString), &profileToFollow)

	if err!=nil{
		http.Error(rw, "Can't unmarshal body!", http.StatusBadRequest)
		return
	}

	fmt.Println("********** Treba da nam postane follower: ",profileToFollow.FollowToUsername)

	myProfile, errNotFound := u.Service.GetProfileByUsername(meDto.Username)

	if errNotFound!=nil{
		fmt.Println("Not Found: ", meDto.Username)
		http.Error(rw, "Not Found my profile", http.StatusBadRequest)
		return
	}

	profileForFollow, errNotFound := u.Service.GetProfileByUsername(profileToFollow.FollowToUsername)
	if errNotFound!=nil{
		fmt.Println("Not Found: ", profileToFollow.FollowToUsername)
		http.Error(rw, "Not Found follower's profile", http.StatusBadRequest)
	}

	fmt.Println("JA SAM: ", myProfile.Username)
	fmt.Println("FOLLOWER MI POSTAJE: ", profileForFollow.Username)

	err = u.Service.AcceptFollow(myProfile, profileForFollow)

	if err!=nil{
		http.Error(rw, "Error at Accepting profile for follow", http.StatusInternalServerError)
		return
	}

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
	url := "http://" + constants.AUTH_SERVICE_URL + "/whoami"
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error with the whoami request")
	}
	req.Header.Add("Authorization", tokenString)
	return client.Do(req)
}


func getCurrentUserCredentials(tokenString string) (dto.UsernameRole, error) {

	resp, err := UserCheck(tokenString)
	if err != nil {
		//p.L.Fatalln("There has been an error sending the /whoami request")
		//rw.WriteHeader(http.StatusInternalServerError)
		return dto.UsernameRole{}, fmt.Errorf("Error sending who am I request")
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var dto dto.UsernameRole
	err = dto.URFromJSON(resp.Body)

	if err != nil {
		fmt.Errorf("Error in unmarshaling JSON")
	}

	return dto, nil

}

func (handler *ProfileHandler) GetCurrent(rw http.ResponseWriter, r *http.Request) {
	// send whoami to auth service
	dto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	fmt.Println(dto.Username + "----" + dto.Role)
	rw.Header().Set("Content-Type", "application/json")
	profile := handler.Service.GetCurrentProfile(dto.Username)

	profile.ToJson(rw)

	rw.WriteHeader(http.StatusOK)

}

func (u *ProfileHandler) GetUser(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	username := params["username"]

	profile, err := u.Service.GetProfileByUsername(username)
	if err != nil {
		http.Error(
			writer,
			fmt.Sprintf("Error fetching user from repo%s", err),
			http.StatusInternalServerError,
		)
		return
	}
	profile.ToJson(writer)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}


func (handler *ProfileHandler) BlockUser(rw http.ResponseWriter, r *http.Request) {
	var blockdto dto.BlockmuteDTO

	// send whoami to auth service
	dto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	fmt.Println("JA SAM : " + dto.Username + "----" + dto.Role)
	rw.Header().Set("Content-Type", "application/json")

	error := blockdto.FromJSON(r.Body)

	if error != nil {
		fmt.Errorf("Error in unmarshaling blockdto")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	profile, err := handler.Service.BlockProfile(dto.Username, blockdto.UsernameToBlockMute)

	if profile == nil {
		fmt.Errorf("Error in blockprofile")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("AFTER SAVE PROFILE : ")
	fmt.Println(profile)

	fmt.Print("I AM //" + dto.Username + "// I BLOCK: //" + blockdto.UsernameToBlockMute + "//")

	blockdto.ToJSON(rw)

	rw.WriteHeader(http.StatusOK)
}


func (handler *ProfileHandler) MuteUser(rw http.ResponseWriter, r *http.Request) {
	var mutedto dto.BlockmuteDTO

	// send whoami to auth service
	dto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	fmt.Println("JA SAM : " + dto.Username + "----" + dto.Role)
	rw.Header().Set("Content-Type", "application/json")

	error := mutedto.FromJSON(r.Body)

	if error != nil {
		fmt.Errorf("Error in unmarshaling blockdto")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("USER TO BLOCK: " + mutedto.UsernameToBlockMute)
	profile, err := handler.Service.MuteProfile(dto.Username, mutedto.UsernameToBlockMute)

	if profile == nil {
		fmt.Errorf("Error in muteprofile")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("AFTER SAVE PROFILE : ")
	fmt.Println(profile)

	fmt.Print("I AM //" + dto.Username + "// I BLOCK: //" + mutedto.UsernameToBlockMute + "//")

	mutedto.ToJSON(rw)

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) UnblockUser(rw http.ResponseWriter, r *http.Request) {
	var blockdto dto.BlockmuteDTO

	// send whoami to auth service
	dto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	fmt.Println(dto)
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	error := blockdto.FromJSON(r.Body)

	if error != nil {
		fmt.Errorf("Error in unmarshaling blockdto")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	profile, err := handler.Service.UnblockProfile(dto.Username, blockdto.UsernameToBlockMute)

	if profile == nil {
		fmt.Errorf("Error in unblockprofile")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	blockdto.ToJSON(rw)

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) Unmute(rw http.ResponseWriter, r *http.Request) {
	var mutedto dto.BlockmuteDTO

	// send whoami to auth service
	dto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	error := mutedto.FromJSON(r.Body)

	if error != nil {
		fmt.Errorf("Error in unmarshaling blockdto")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	profile, err := handler.Service.UnmuteProfile(dto.Username, mutedto.UsernameToBlockMute)

	if profile == nil {
		fmt.Errorf("Error in unblockprofile")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	mutedto.ToJSON(rw)

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) AddPostToFavourites(rw http.ResponseWriter, r *http.Request) {
	var postToFavourites dto.PostToFavourites

	userRoleDto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	err = postToFavourites.FromJSON(r.Body)

	if err != nil {
		fmt.Errorf("Error in unmarshaling postToFavourites")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.AddPostToFavourites(userRoleDto, postToFavourites)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf(err.Error()),
			http.StatusBadRequest,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) GetFavourites(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	collection := params["collection"]
	fmt.Println(collection)

	userRoleDto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	postsIds, err := handler.Service.GetPostsIdsInCollection(collection, userRoleDto)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error getting posts IDS from repository", err),
			http.StatusInternalServerError,
		)
		return
	}

	for _, p := range postsIds.Ids {
		fmt.Println("POST:")
		fmt.Println(p)
	}

	err = postsIds.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to unmarshal posts json", http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) GetAllPublicProfiles(rw http.ResponseWriter, r *http.Request) {

	err, profiles := handler.Service.GetAllPublicProfiles()

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf(err.Error()),
			http.StatusInternalServerError,
		)
	}
	err = profiles.ToJson(rw)
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf(err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
// gets all profiles that are private but this user follows them
func (handler *ProfileHandler) GetAllowedProfiles(rw http.ResponseWriter, r *http.Request) {

	dto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}
	profiles := handler.Service.GetAllAllowedProfiles(dto.Username)

	err = profiles.ToJson(rw)
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error serializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) GetAllNonFollowedPrivateProfiles(rw http.ResponseWriter, r *http.Request) {

	dto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	err, profiles := handler.Service.GetAllNonFollowedPrivateProfiles(dto.Username)

	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error fetching private non followed users  %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	profiles.ToJson(rw)
	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) GetUsersWhoBlockedMe(rw http.ResponseWriter, r *http.Request) {
	var retlist dto.ListWhoBlockedMeDTO
	dto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	err ,usersListWhoBlockedMe := handler.Service.GetUsersWhoBlockedMe(dto.Username)

	fmt.Print("------------------------------------")
	fmt.Println(len(usersListWhoBlockedMe))

	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error getting users who blocked me"),
			http.StatusInternalServerError,
		)
		return
	}

	retlist.ListWhoBlockedMe = usersListWhoBlockedMe

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error marshaling usersListWhoBlockedMe"),
			http.StatusInternalServerError,
		)
		return
	}
	rw.Header().Set("Content-Type", "application/json")

	retlist.ToJSON(rw)

	rw.WriteHeader(http.StatusOK)

}

func (handler *ProfileHandler) GetAllFavourites(rw http.ResponseWriter, r *http.Request) {
	userRoleDto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	collections, err := handler.Service.GetFavouritPostsIds(userRoleDto)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error getting posts IDS from repository", err),
			http.StatusInternalServerError,
		)
		return
	}

	err = collections.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal posts json", http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) DeleteCollection(rw http.ResponseWriter, r *http.Request) {
	var postToFavourites dto.PostToFavourites
	userRoleDto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	err = postToFavourites.FromJSON(r.Body)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error in unmarshaling postToFavourites"),
			http.StatusInternalServerError,
		)
		return
	}

	err = handler.Service.DeleteCollection(userRoleDto.Username, postToFavourites.CollectionName)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error in deleting collection"),
			http.StatusInternalServerError,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) DeletePostFromCollection(rw http.ResponseWriter, r *http.Request) {
	var postToFavourites dto.PostToFavourites
	userRoleDto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusBadRequest,
		)
		return
	}

	err = postToFavourites.FromJSON(r.Body)

	fmt.Println(postToFavourites)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error in unmarshaling postToFavourites"),
			http.StatusInternalServerError,
		)
		return
	}

	err = handler.Service.DeletePostFromCollection(userRoleDto.Username, postToFavourites)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf(err.Error()),
			http.StatusBadRequest,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (handler *ProfileHandler) GetMyNotificationsSettings(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	notifSettings, err := handler.Service.GetMyNotificationsSettings(username)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error: %s", err),
			http.StatusBadRequest,
		)
		return
	}

	//fmt.Println(notifSettings)

	err = notifSettings.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal posts json", http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusOK)
}

func (ph *ProfileHandler) GetCloseFriendsByUsername(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	closeFriends, err := ph.Service.GetCloseFriendsByUsername(username)

	if err!=nil{
		http.Error(rw, "Cant find close friend for passed username", http.StatusNotFound)
		return
	}

	closeFriendsJson, err := json.Marshal(closeFriends)

	if err!=nil{
		http.Error(rw, "Cant Marshal!", http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(closeFriendsJson)
	if err != nil {
		http.Error(rw, "Cant Marshal!", http.StatusInternalServerError)
		return
	}


}

func (u *ProfileHandler) AddProfileToCloseFriends(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameForAddToCloseFriends := params["username"]

	dtoMe, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	err = u.Service.AddProfileToCloseFriends(dtoMe.Username, usernameForAddToCloseFriends)

	if err!=nil{
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)



}

func (u *ProfileHandler) RemoveProfileFromCloseFriends(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameForRemoveFromCloseFriends := params["username"]

	dtoMe, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	err = u.Service.RemoveProfileFromCloseFriends(dtoMe.Username, usernameForRemoveFromCloseFriends)

	if err!=nil{
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)

}

func (u *ProfileHandler) CheckIfCloseFriends(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usernameForCheck := params["username"]

	dtoMe, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}
	fmt.Println(dtoMe.Username)
	fmt.Println(usernameForCheck)
	if dtoMe.Username == usernameForCheck{
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	flag, err := u.Service.CheckIfCloseFriends(dtoMe.Username, usernameForCheck)

	if err!=nil{
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if flag==true{
		rw.WriteHeader(http.StatusAccepted)
		return
	}


	rw.WriteHeader(http.StatusOK)


}



func GetCurrentUserWrapper(tokenString string) (dto.UsernameRole, error) {

	return getCurrentUserCredentials(tokenString)
}
