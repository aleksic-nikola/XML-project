package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"xml/auth-service/data"
	"xml/auth-service/dto"
	"xml/auth-service/security"
	"xml/auth-service/service"

	"github.com/dgrijalva/jwt-go"
)

type UserHandler struct {
	L       *log.Logger
	Service *service.UserService
}

func NewUsers(l *log.Logger, service *service.UserService) *UserHandler {
	return &UserHandler{l, service}
}

// Login function
// returns 400 if any error occurs
// returns 401 if credentials are invalid
// returns 200 with token if successfull
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
	rw.Header().Set("Authorization", "Bearer "+token)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Token: " + token))
}

// deserializes the body object into a json
// throws 400 on any kind of error
// hashes password
// saves it to db
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
		http.Error(rw, "Unable to unmarshal users json", http.StatusInternalServerError)
	}
}

// catches any request to a certain URL and cuts it off
// checks the authorization header disects the token and sends it back as a header parameter
// the actual handler function should look at the authorization and decide whether it's allowed or not
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := security.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		username := claims.(jwt.MapClaims)["username"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("username", username)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}


func (handler *UserHandler) EditUserData(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("updating user data")
	var dto dto.UserEditDTO
	err := dto.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dto)

	// uzeti username od trenutno ulogovanog korisnika - auth
	// proslediti to u service
	oldUsername := "wintzyboi"
	err = handler.Service.EditUserData(dto, oldUsername)

	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
}
