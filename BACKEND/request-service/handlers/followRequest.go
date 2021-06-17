package handlers

import (
	"bytes"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"xml/request-service/data"
	dtoRequest "xml/request-service/dto"
	"xml/request-service/service"
)

type FollowRequestHandler struct {
	L *log.Logger
	Service *service.FollowRequestService
}

func NewFollowRequest(l *log.Logger, service *service.FollowRequestService) *FollowRequestHandler {
	return &FollowRequestHandler{l, service}
}

func (handler *FollowRequestHandler) CreateFollowRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating...")
	var followReq data.FollowRequest
	err := followReq.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.CreateFollowRequest(&followReq)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	fmt.Println("Created")
	rw.Header().Set("Content-Type", "application/json")
}



func BodyToJson(body io.ReadCloser) (string, error) {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return "", fmt.Errorf("Error with reading body...")
	}
	return string(bodyBytes), nil
}




func (p *FollowRequestHandler) GetFollowRequests(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request")

	lp := data.GetFollowRequests()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}


func (p *FollowRequestHandler) GetMyFollowRequests(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request")

	jwtToken := r.Header.Get("Authorization")
	resp, err := UserCheck(jwtToken)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonString, err := BodyToJson(resp.Body)

	if err!=nil{
		fmt.Println("Error BodyToJson...")
		return
	}

	var meDto dtoRequest.UsernameRoleDto
	err = json.Unmarshal([]byte(jsonString), &meDto)
	if err != nil {
		fmt.Println("Error at Unmarshal")
		return
	}

	myFollReqs, err := p.Service.GetMyFollowRequests(meDto.Username)

	myFollReqsJson, err := json.Marshal(myFollReqs)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	_, err1 := rw.Write(myFollReqsJson)

	if err1 != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
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



func (p *FollowRequestHandler) AcceptFollowRequest(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle POST Request")

	jsonString, err := BodyToJson(r.Body)
	if err!=nil{
		fmt.Println("Error BodyToJson...")
		return
	}

	var followReq dtoRequest.FollowRequestDto
	err = json.Unmarshal([]byte(jsonString), &followReq)
	if err != nil {
		fmt.Println("Error at Unmarshal")
		return
	}

	fmt.Println("OD: ", followReq.SentBy)
	fmt.Println("ZA: (MENE): ", followReq.ForWho)

	if followReq.SentBy == followReq.ForWho{
		http.Error(rw, "Error request!", http.StatusBadRequest)
		return
	}

	jwtToken := r.Header.Get("Authorization")
	client := &http.Client{}
	var dtoUsername dtoRequest.ProfileForFollow

	dtoUsername.FollowToUsername = followReq.SentBy

	usernameJson, err := json.Marshal(dtoUsername)

	url := "http://localhost:3030/acceptFollow"//-------------> Adding new profile to my Followers and me to his Following
	req, errReq := http.NewRequest("POST", url, bytes.NewBuffer(usernameJson))

	if errReq != nil{
		http.Error(rw, "Cant accept request", http.StatusBadRequest)
	}
	req.Header.Add("Authorization", jwtToken)
	resp, err := client.Do(req)

	if err != nil{
		http.Error(rw, "Error with sending request...", http.StatusBadRequest)
		return
	}

	if resp.StatusCode!=200 {
		fmt.Println(resp.StatusCode)
		http.Error(rw, "Cant accept follow --> profileService", http.StatusInternalServerError)
		return
	}

	err = p.Service.AcceptFollowRequest(followReq.SentBy, followReq.ForWho) //change RequestStatus to ACCEPTED
	if err != nil {
		fmt.Println("Can't find req")
		return
	}
	rw.WriteHeader(http.StatusOK)
}
