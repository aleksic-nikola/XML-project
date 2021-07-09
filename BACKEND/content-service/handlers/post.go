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
	"strconv"
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
	fmt.Println("TOKEN: " + r.Header.Get("Authorization"))
	dto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	userID, err := GetIDByUsername(dto.Username, r.Header.Get("Authorization"))
	if err!=nil{
		fmt.Println("ERROR GETTING ID.......")
		return
	}

	err = r.ParseMultipartForm(2000000) // grab the multipart form
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
	desc := res["description_part"]
	fmt.Println(res["description_part"])
	fmt.Println(id)

	path, _ := filepath.Abs("./")
	fmt.Println(filepath.Join(path, "temp")	)
	//tempFile, err := ioutil.TempFile(filepath.Join(path, "temp"), "upload-*.png")
	fmt.Println(filepath.Join(path, "temp/id-" + strconv.Itoa(int(userID))))
    //tempFile, err := ioutil.TempFile(filepath.Join(path, "temp"), "upload-*.png")
	var medias []data.Media

	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(rw, err)
			return
		}
		//prepath, err:= os.Getwd()
		//finalpath := filepath.Join(prepath, "temp")
		finalPath := ""
		if os.Getenv("DOCKERIZED")== "yes" {
			fmt.Println("USAO DA JE DOKERIZED!")
			finalPath = "./temp/id-" + strconv.Itoa(int(userID)) + "/" + files[i].Filename
		} else{
			finalPath =  filepath.Join("../../FRONTEND/frontend-service/static/temp/id-" + strconv.Itoa(int(userID)), files[i].Filename )

		}

		out, err := os.Create(finalPath)

		fmt.Println("final path is" + finalPath)

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
		// TODO: FIX THIS
		var media data.Media
		media.Path = "temp/id-" + strconv.Itoa(int(userID)) + "/" + files[i].Filename
		var location data.Location
		location.Country = "some country"
		location.City = "some_city"
		media.Location = location
		media.Type = "IMAGE"
		medias = append(medias, media)

		//fmt.Fprintf(rw, "Files uploaded successfully : ")
		//fmt.Fprintf(rw, files[i].Filename+"\n")

	}
	// for is over
	// create the actual post with uris

	var post data.Post
	post.Medias = medias
	post.Description = desc[0]
	post.PostedBy = dto.Username

	err = p.Service.CreatePost(&post)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error creating POST %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)




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

func GetIDByUsername(username string, tokenString string) (uint, error){
	godotenv.Load()
	client := &http.Client{}
	url := "http://" + constants.AUTH_SERVICE_URL + "/getuserId/" + username
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("Error with the whoami request")
	}
	req.Header.Add("Authorization", tokenString)
	resp, err := client.Do(req)

	if err!=nil{
		fmt.Println("GRESKA ZAHTEV!!!")
		return 0, err
	}

	var userId dtos.UserIDDto

	err = userId.FromJSON(resp.Body)
	if err!=nil{
		fmt.Println("Error unmarshaling!!!!")
		return 0, err
	}
	return uint(userId.UserId), nil

}




func getCurrentUserCredentials(tokenString string) (dtos.UsernameRole, error) {
	fmt.Println("OVO JE TOKEN: ")
	fmt.Println(tokenString)
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
	err = dto.FromJSON(resp.Body)
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

func (handler *PostHandler) CreateDirectoryForUser(rw http.ResponseWriter, r *http.Request) {
	var userID dtos.UserIDDto
	err := userID.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	path, _ := filepath.Abs("./")
	fmt.Println("************* PUTANJA **********************")
	fmt.Println(filepath.Join(path, "temp" + strconv.Itoa(userID.UserId)))
	wd, err := os.Getwd() // content-serivce
	parent := filepath.Dir(wd) // BACKEND
	parent = filepath.Dir(parent) // root
	make_folder_here := filepath.Join(parent, "FRONTEND", "frontend-service", "temp/id-" + strconv.Itoa(userID.UserId))
	fmt.Println(make_folder_here)
	fmt.Println("*************")
	/*
	err = os.Mkdir(make_folder_here, 0755)
	if err !=nil{
		fmt.Println("Error at creating directory")
		return
	}
	*/
	rw.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) LikePost(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	post_id := params["post"]

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


	err = handler.Service.LikePost(post_id, myUsernameRoleDto.Username)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error liking post %s", err),
			http.StatusBadRequest,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) DislikePost(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	post_id := params["post"]

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


	err = handler.Service.DislikePost(post_id, myUsernameRoleDto.Username)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error disliking post %s", err),
			http.StatusBadRequest,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) Comment(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	post_id := params["post"]

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

	var dto dtos.CommentDto
	err = dto.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.PostComment(post_id, dto.Text, myUsernameRoleDto.Username)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error posting comment %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) GetFavouritePosts(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	collection := params["collection"]
	fmt.Println(collection)

	userRoleDto, err := getCurrentUserCredentials(r.Header.Get("Authorization"))
	if err != nil {

		http.Error(
			rw,
			fmt.Sprintf("Error deserializing JSON %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	fmt.Println(userRoleDto)

	var resp, err1 = GetFavouritePostsIds(r.Header.Get("Authorization"), collection)

	if err1 != nil{
		fmt.Println("Respond error!!!")
		http.Error(rw, "Respond error getPOSTIDS!!!", http.StatusInternalServerError)
		return
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

	var postIds dtos.PostIdsDto

	err = json.Unmarshal(b, &postIds)
	if err !=nil{
		fmt.Println("Error at unmarsal allFollowing")
		return
	}
	fmt.Println("DOBIO FROMJSON: ")
	fmt.Println(postIds)

	posts := handler.Service.GetPostsByIds(postIds)

	err = posts.ToJSON(rw)

	rw.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) GetUserByPost(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postid := params["id"]
	fmt.Println(postid)

	username := handler.Service.GetUserByPost(postid)

	err := username.ToJSON(rw)

	if err !=nil{
		fmt.Println("Error at ToJSONING username")
		return
	}

	rw.WriteHeader(http.StatusOK)
}


func GetFavouritePostsIds(tokenString string, collection string) (*http.Response, error) {
	godotenv.Load()

	client := &http.Client{}
	url := "http://" + constants.PROFILE_SERVICE_URL + "/getFavouritePosts/" + collection
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("Error with creating new verified user")
	}
	req.Header.Add("Authorization", tokenString)
	return client.Do(req)
}
