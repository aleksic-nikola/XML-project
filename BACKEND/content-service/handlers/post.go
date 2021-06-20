package handlers

import (
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

