package handlers

import (
	"fmt"
	"log"
	"net/http"
	
	"xml/content-service/data"
	"xml/content-service/data/dtos"
	"xml/content-service/service"

	"github.com/joho/godotenv"
)

type PostHandler struct {
	L       *log.Logger
	Service *service.PostService
}

func NewPosts(l *log.Logger, service *service.PostService) *PostHandler {
	return &PostHandler{l, service}
}

func (handler *PostHandler) CreatePost(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var post data.Post
	err := post.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(post)

	err = handler.Service.CreatePost(&post)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (p *PostHandler) GetPosts(rw http.ResponseWriter, r *http.Request) {
	p.L.Println("Handle GET Request for Posts")

	lp := data.GetPosts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal posts json", http.StatusInternalServerError)
	}
}

func (p *PostHandler) GetPostsForCurrentUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	// send whoami to auth service
	godotenv.Load()
	client := &http.Client{}
	tokenString := r.Header.Get("Authorization")
	url := "http://" + GetVariable("auth") + "/whoami"
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	fmt.Println("here")
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	//var bearer = "Bearer " + tokenString
	req.Header.Add("Authorization", tokenString)
	resp, err := client.Do(req)
	if err != nil {
		p.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Println("here")
	fmt.Println(resp.Status)
	var dto dtos.UsernameRole
	err = dto.FromJSON(resp.Body)
	fmt.Println("---------------")
	fmt.Println(resp.Body)
	if err != nil {
		
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}
	fmt.Println(dto)
	dto.ToJSON(rw)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	//resp, err := http.Get(os.Getenv("profile") + "/whoami")
}
