package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"xml/content-service/constants"
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

func (handler *PostHandler) GetPostsByUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	lp := handler.Service.GetPostsByUser(username)
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal posts json", http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusOK)

}

func (p *PostHandler) GetPostsForCurrentUser(rw http.ResponseWriter, r *http.Request) {
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
	posts := p.Service.GetAllPostsForUser(dto.Username)
	err = posts.ToJSON(rw)
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,	
		)
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	//resp, err := http.Get(os.Getenv("profile") + "/whoami")
}

func (p *PostHandler) UploadPost(rw http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(200000) // grab the multipart form
 	if err != nil {
 		fmt.Fprintln(rw, err)
 		return
 	}
	 formdata := r.MultipartForm // ok, no problem so far, read the Form data
	//get the *fileheaders
	files := formdata.File["multiplefiles"] // grab the filenames
	fmt.Println(formdata.Value)
	res := formdata.Value
	id  := res["post_title"]
	fmt.Println(res["description_part"])
	fmt.Println(id)

	path, _ := filepath.Abs("./") 
    	fmt.Println(filepath.Join(path, "temp")	)
    	//tempFile, err := ioutil.TempFile(filepath.Join(path, "temp"), "upload-*.png")
	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(rw, err)
			return
		}

		out, err := os.Create(filepath.Join(path, "temp", files[i].Filename))

		defer out.Close()
		if err != nil {
			fmt.Fprintf(rw, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			fmt.Fprintln(rw, err)
			return
		}

		fmt.Fprintf(rw, "Files uploaded successfully : ")
		fmt.Fprintf(rw, files[i].Filename+"\n")

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

func getCurrentUserCredentials(tokenString string) (dtos.UsernameRole, error) {

	resp, err := UserCheck(tokenString)
	if err != nil {
		//p.L.Fatalln("There has been an error sending the /whoami request")
		//rw.WriteHeader(http.StatusInternalServerError)
		return dtos.UsernameRole{}, fmt.Errorf("Error sending who am I request")
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var dto dtos.UsernameRole
	err = dto.FromJSON(resp.Body)
	if err != nil {
		fmt.Errorf("Error unmarshalling JSON from response body")
	}
	return dto, nil
}

func (p *PostHandler) GetLikedPostsByUser(rw http.ResponseWriter, r *http.Request) {
	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		p.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var dto dtos.UsernameRole
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	posts := p.Service.GetLikedPostsByUser(dto.Username)
	err = posts.ToJSON(rw)
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}

func (p *PostHandler) GetDislikedPostsByUser(rw http.ResponseWriter, r *http.Request) {
	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		p.L.Fatalln("There has been an error sending the /whoami request")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var dto dtos.UsernameRole
	err = dto.FromJSON(resp.Body)
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	posts := p.Service.GetDislikedPostsByUser(dto.Username)
	err = posts.ToJSON(rw)
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

}

func (h *PostHandler) GetAllPostsForFeedForCurrentUser(rw http.ResponseWriter, r *http.Request) {
	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		//p.L.Fatalln("There has been an error sending the /whoami request")
		http.Error(rw,"There has been an error sending the /whoami request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	var myUsernameRoleDto dtos.UsernameRole
	err = myUsernameRoleDto.FromJSON(resp.Body)
	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	//pronaci sve Following

	godotenv.Load()
	client := &http.Client{}
	url := "http://" + constants.PROFILE_SERVICE_URL + "/getAllFollowing"
	fmt.Println(url)

	var usernameDto dtos.UsernameDto
	usernameDto.Username = myUsernameRoleDto.Username

	usernameDtoJson, err := json.Marshal(usernameDto)

	if err!= nil{
		fmt.Println("Unable to marsh usernameDto")
		http.Error(rw, "Unable to marsh usernameDto", http.StatusBadRequest)
		return
	}



	req, err := http.NewRequest("POST", url, bytes.NewBuffer(usernameDtoJson))
	if err != nil {
		//return nil, fmt.Errorf("Error with the whoami request")
	}
	req.Header.Add("Authorization", r.Header.Get("Authorization"))

	resp, err = client.Do(req)
	fmt.Println(resp.Body)

	if err!=nil{
		fmt.Println("Respond error!!!")
		http.Error(rw, "Respond error getFollowing!!!", http.StatusInternalServerError)
		return
	}


	fmt.Println("DOBILI: ")
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

	//allUsers := make([]dtos.Profile, 0)

	var allUsers []dtos.UsernameDto
	err1 := json.Unmarshal(b, &allUsers)
	if err1!=nil{
		fmt.Println("Error at unmarsal allFollowing")
		return
	}
	fmt.Println("DOBIO FROMJSON: ")
	fmt.Println(allUsers)

	var allUsernames []string

	for _, oneUser:= range allUsers{
		allUsernames = append(allUsernames, oneUser.Username)
	}
	allUsernames = append(allUsernames, myUsernameRoleDto.Username)
	fmt.Println(allUsernames)

	allPosts := h.getAllPostsForFeed(allUsernames)

	fmt.Println(allPosts)

	allPostsJson, _ := json.Marshal(allPosts)

	_, err = rw.Write(allPostsJson)
	if err!=nil{
		fmt.Println("Error with Write!")
		return
	}

}

func (h *PostHandler) getAllPostsForFeed(usernames []string) []data.Post {
	allPosts := h.Service.GetAllPostsForFeed(usernames)

	return allPosts
}

