package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"strconv"
	"xml/interaction-service/constants"
	"xml/interaction-service/data"
	"xml/interaction-service/dto"
	"xml/interaction-service/service"
)

type PostNotificationHandler struct {
	L       *log.Logger
	Service *service.PostNotificationService
}

func NewPostNotifications(l *log.Logger, service *service.PostNotificationService) *PostNotificationHandler {
	return &PostNotificationHandler{l, service}
}

func (handler *PostNotificationHandler) CreatePostNotification(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")

	//Getting current user(*notification from)
	userdto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error: %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	var notifdto dto.PostNotif
	err = notifdto.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(notifdto)

	//getting user who posted (*notification for)
	var resp, err1 = GetUserByPostID(notifdto.Post_id)

	if err1 != nil{
		fmt.Println("Respond error!!!")
		http.Error(rw, "Respond error getPOSTIDS!!!", http.StatusInternalServerError)
		return
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

	var sentBy dto.UsernameDto

	err = json.Unmarshal(b, &sentBy)
	if err !=nil{
		fmt.Println("Error at unmarsal allFollowing")
		return
	}

	notificationSettingDto, err := getNotificationSettings(sentBy.Username)
	fmt.Println(notificationSettingDto)

	//PROVERE DA LI je sam sebi nesto uradio
	if sentBy.Username == userdto.Username {
		rw.WriteHeader(http.StatusOK)
		return
	}

	//PROVERE DA LI USER ZELI DA DOBIJA NOTIFIKACIJE
	if notifdto.Type == "COMMENT" {
		if notificationSettingDto.ShowDmNotification == false {
			rw.WriteHeader(http.StatusOK)
			return
		}

	} else {
		if notificationSettingDto.ShowTaggedNotification == false {
			rw.WriteHeader(http.StatusOK)
			return
		}
	}


	err2 := handler.Service.CreatePostNotification(notifdto, userdto.Username, sentBy.Username)


	if err2 != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func GetUserByPostID(postId int) (*http.Response, error) {
	godotenv.Load()

	client := &http.Client{}

	id := strconv.Itoa(postId)
	fmt.Println(id)
	url := "http://" + constants.CONTENT_SERVICE_URL + "/getuserbypostid/" + id
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("Error getting user by postid")
	}
	return client.Do(req)
}

func (p *PostNotificationHandler) GetPostNotifications(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request in postNotificationHandler")

	lp := data.GetPostNotifications()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (handler *PostNotificationHandler) GetUnreadPostNotif(rw http.ResponseWriter, r *http.Request) {
	userdto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}
	fmt.Println(userdto.Username)

	notifications := handler.Service.GetMyUnreadPostNotif(userdto.Username)

	err = notifications.ToJSON(rw)
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
}

func (handler *PostNotificationHandler) ReadPostNotifications(rw http.ResponseWriter, r *http.Request) {
	userdto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}
	fmt.Println(userdto.Username)

	err = handler.Service.ReadPostNotifications(userdto.Username)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error changing read status %s", err),
			http.StatusBadRequest,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)
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
	err = dto.FromJSON(resp.Body)

	if err != nil {
		fmt.Errorf("Error in unmarshaling JSON")
	}

	return dto, nil

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

func getNotificationSettings(username string) (dto.NotificationSetting, error) {
	var resp, err1 = GetNotificationSetting(username)

	if err1 != nil{
		fmt.Println("Respond error!!!")
		//http.Error(rw, "Respond error getPOSTIDS!!!", http.StatusInternalServerError)
		return dto.NotificationSetting{}, nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
	log.Fatalln(err)
	}
	fmt.Println(string(b))

	var notificationSetting dto.NotificationSetting

	err = json.Unmarshal(b, &notificationSetting)
	if err !=nil {
		fmt.Println("Error at unmarsal allFollowing")
		return dto.NotificationSetting{}, err
	}

	return notificationSetting, nil
}

func GetNotificationSetting(username string) (*http.Response, error) {
	godotenv.Load()

	client := &http.Client{}

	url := "http://" + constants.PROFILE_SERVICE_URL + "/getmynotifsettings/" + username
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("Error getting user by postid")
	}
	return client.Do(req)
}