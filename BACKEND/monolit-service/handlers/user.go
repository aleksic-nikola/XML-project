package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/monolit-service/data"
	"xml/monolit-service/service"
)

type UserHandler struct {
	L *log.Logger
	Service *service.UserService
}

func NewUsers(l *log.Logger, service *service.UserService) *UserHandler {
	return &UserHandler{l, service}
}

func (handler *UserHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating user")
	var user data.User
	err := user.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)

	err = handler.Service.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (u *UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	u.L.Println("Handle GET Request for Users")

	lp := data.GetUsers()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal users json" , http.StatusInternalServerError)
	}
}