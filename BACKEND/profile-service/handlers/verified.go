package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/profile-service/data"
	"xml/profile-service/dto"
	"xml/profile-service/service"
)

type VerifiedHandler struct {
	L *log.Logger
	Service *service.VerifiedService
	ProfileService *service.ProfileService
}

func NewVerifieds(l *log.Logger, service *service.VerifiedService, profileService *service.ProfileService) *VerifiedHandler {
	return &VerifiedHandler{l, service, profileService}
}

func (u *VerifiedHandler) GetVerifieds(rw http.ResponseWriter, r *http.Request) {
	u.L.Println("Handle GET Request for Profiles")

	lp := data.GetVerifieds()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}

func (handler *VerifiedHandler) CreateVerified(rw http.ResponseWriter, r *http.Request) {
	handler.L.Println("Create verified starting")

	var newVerified dto.NewVerified

	err := newVerified.FromJSON(r.Body)

	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(newVerified)
	profile, err := handler.ProfileService.GetProfileByUsername(newVerified.Username)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error fetching user from repo%s", err),
			http.StatusInternalServerError,
		)
		return
	}

	err = handler.Service.CreateNewVerified(profile, newVerified.VerifiedType)

	if err != nil {
		fmt.Println("Error creating verified profile")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *VerifiedHandler) CheckIfVerified(rw http.ResponseWriter, r *http.Request) {

	userRoleDto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))

	if err != nil {
		fmt.Println("Error getting verified status")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("><><><><><")
	fmt.Println(userRoleDto)
	id,err := u.ProfileService.GetIdByUsername(userRoleDto.Username)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf(err.Error()),
			http.StatusBadRequest,
		)
		return
	}

	ver, err1 := u.Service.GetVerificationForUser(id)
	fmt.Println("CheckIfVerified Error")
	fmt.Println(err1)
	if err1 != nil {
		fmt.Println("THIS USER IS NOT VERIFIED BRO")
		http.Error(
			rw,
			fmt.Sprintf(err1.Error()),
			http.StatusBadRequest,
		)
		return
	}

	if ver.Profile.Username == "" {
		fmt.Println("THIS USER IS NOT VERIFIED BRO")
		http.Error(
			rw,
			fmt.Sprintf("This user is not verified"),
			http.StatusBadRequest,
		)
		return
	}
	ver.ToJson(rw)
	rw.WriteHeader(http.StatusOK)
}

