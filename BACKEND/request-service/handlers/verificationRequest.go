package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"xml/request-service/constants"
	"xml/request-service/data"
	dtoRequest "xml/request-service/dto"
	"xml/request-service/service"
)

type VerificationRequestHandler struct {
	L *log.Logger
	Service *service.VerificationRequestService

}

func NewVerificationRequest(l *log.Logger, service *service.VerificationRequestService) *VerificationRequestHandler {
	return &VerificationRequestHandler{l, service}
}


func (handler *VerificationRequestHandler) CreateVerificationRequest(rw http.ResponseWriter, r *http.Request) {
	var verificationRequest data.VerificationRequest
	err := verificationRequest.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(verificationRequest)

	err = handler.Service.CreateVerificationRequest(&verificationRequest)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}


func (handler *VerificationRequestHandler) GetInProgressVerificationRequests(rw http.ResponseWriter, r *http.Request) {
	handler.L.Println("Handle GET Request")

	verificationRequests := handler.Service.GetInProgressVerificationRequests()

	err := verificationRequests.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusOK)

}

func (handler *VerificationRequestHandler) CreateInProgressVerificationRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating new VerificationRequest")
	var verificationReqDto dtoRequest.VerificationRequestDto

	r.ParseMultipartForm(2000000) // grab the multipart form

	formdata := r.MultipartForm // ok, no problem so far, read the Form data
	//get the *fileheaders

	fmt.Println(formdata.Value)
	res := formdata.Value
	name := res["verification_name"]
	lastname := res["verification_lastname"]
	category := res["verification_category"]

	fmt.Println(name)
	fmt.Println(lastname)
	fmt.Println(category)

	fmt.Println(verificationReqDto)

	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		handler.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var dto dtoRequest.UsernameRoleDto
	err = dto.FromJSON(resp.Body)
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	client := &http.Client{}
	url := "http://" + constants.AUTH_SERVICE_URL + "/getuserId/" + dto.Username
	req1, errReq := http.NewRequest("GET", url, nil)
	fmt.Println(url)
	if errReq != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error with who am I request %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	req1.Header.Add("Authorization", r.Header.Get("Authorization"))

	resp1, err := client.Do(req1)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error getting user Id by username %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	defer resp1.Body.Close()
	var user_id dtoRequest.UserIdDto
	err = user_id.FromJSON(resp1.Body)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error unmarshalling user ID JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}
	// file handling

	file, h, err := r.FormFile("filephoto")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", h.Filename)
	fmt.Printf("File Size: %+v\n", h.Size)
	fmt.Printf("MIME Header: %+v\n", h.Header)

	fileForRequest, err := h.Open()
	defer fileForRequest.Close()
	if err != nil {
		fmt.Fprintln(rw, err)
		return
	}
	var finalPath string
	if os.Getenv("DOCKERIZED")== "yes" {
		fmt.Println("USAO DA JE DOKERIZED!")
		finalPath = "./temp/id-" + strconv.Itoa(int(user_id.UserId)) + "/" + h.Filename
		fmt.Println(finalPath)
	} else{
		finalPath =  filepath.Join("../../FRONTEND/frontend-service/static/temp/id-" + strconv.Itoa(int(user_id.UserId)), h.Filename )
	}

	out, err := os.Create(finalPath)

	if err != nil {
		fmt.Errorf("Error in creating")
		fmt.Println("error in creating")
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}

	_, err = io.Copy(out, file) // file not files[i] !

	if err != nil {
		fmt.Fprintln(rw, err)
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}
	var verifiedType data.VerifiedType

	if category[0] == "influencer" {
		verifiedType = 0
	} else if category[0] == "sports" {
		verifiedType = 1
	} else if category[0] == "media" {
		verifiedType = 2
	} else if category[0] == "business" {
		verifiedType = 3
	} else if category[0] == "brand" {
		verifiedType = 4
	} else {
		verifiedType = 5
	}

	verificationReqDto.Name = name[0]
	verificationReqDto.LastName = lastname[0]
	verificationReqDto.Image = "temp/id-" + strconv.Itoa(int(user_id.UserId)) + "/" + h.Filename
	verificationReqDto.Category = verifiedType
	err = handler.Service.CreateInProgressVerificationRequest(&verificationReqDto, &dto)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (handler *VerificationRequestHandler) UpdateVerificationRequest(rw http.ResponseWriter, r *http.Request) {

	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		handler.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var dto dtoRequest.UsernameRoleDto
	err = dto.FromJSON(resp.Body)
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	if dto.Role != "admin" {
		http.Error(
			rw,
			fmt.Sprintf("Only admin can update verification requests", err),
			http.StatusMethodNotAllowed,
		)
		return
	}

	var updateDto dtoRequest.UpdateVerificationRequestDto
	err = updateDto.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(updateDto)

	err1, verificationReq := handler.Service.UpdateVerificationRequest(&updateDto)

	if err1 != nil {
		fmt.Println(err1)
		rw.WriteHeader(http.StatusExpectationFailed)
	}

	if verificationReq.Request.Status == 1 {
		r, err := CreateVerified(*verificationReq)
		fmt.Println(r)
		if err != nil {
			http.Error(
				rw,
				fmt.Sprintf("Error creating verified user", err),
				http.StatusBadRequest,
			)
			return
		}
		rw.WriteHeader(r.StatusCode)
		return
	}
	fmt.Println(verificationReq)

	rw.WriteHeader(http.StatusOK)
}

func CreateVerified(verificationReq data.VerificationRequest) (*http.Response, error) {

	godotenv.Load()
	newVerified := dtoRequest.NewVerified{Username: verificationReq.Request.SentBy,VerifiedType: verificationReq.Category}

	fmt.Println(newVerified)
	requestBody, err := json.Marshal(newVerified)
	client := &http.Client{}
	url := "http://" + constants.PROFILE_SERVICE_URL + "/verified/create"
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("Error with creating new verified user")
	}
	return client.Do(req)
}