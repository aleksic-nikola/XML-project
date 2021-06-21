package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
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
	err := verificationReqDto.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
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

	err = handler.Service.CreateInProgressVerificationRequest(&verificationReqDto, &dto)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
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