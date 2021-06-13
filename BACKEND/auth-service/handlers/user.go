package handlers

import (
	"fmt"
	"log"
	"net/http"
	"xml/auth-service/data"
	"xml/auth-service/security"
	"xml/auth-service/service"
)


type UserHandler struct {
	L *log.Logger
	Service *service.UserService
}

func NewUsers(l *log.Logger, service *service.UserService) *UserHandler {
	return &UserHandler{l, service}
}

func (handler *UserHandler) Login(rw http.ResponseWriter, r *http.Request) {
	var form data.LoginForm
	err := form.LFFromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	user := handler.Service.FindUserByUsername(form.Username)
	if user == nil {
		handler.L.Println("Wrong credentials")
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !security.CheckPasswordHash(form.Password, user.Password) {
		handler.L.Println("Wrong credentials")
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := security.GetToken(user.Username, user.Role)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Authorization", "Bearer " + token)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Token: " + token))
}

func (handler *UserHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var user data.User
	err := user.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	user.Password, err = security.HashPassword(user.Password)
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




