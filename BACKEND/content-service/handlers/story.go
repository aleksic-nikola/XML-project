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
	"xml/content-service/constants"
	"xml/content-service/data"
	"xml/content-service/data/dtos"
	"xml/content-service/service"
)

type StoryHandler struct {
	L       *log.Logger
	Service *service.StoryService
}

func NewStories(l *log.Logger, service *service.StoryService) *StoryHandler {
	return &StoryHandler{l, service}
}

func (handler *StoryHandler) CreateStory(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("creating")
	var story data.Story
	err := story.FromJSON(r.Body)
	if err != nil {
		handler.L.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(story)

	err = handler.Service.CreateStory(&story)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (p *StoryHandler) UploadStory(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("hey im uploading a story bro")
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
		prepath, err:= os.Getwd()
		finalpath := filepath.Join(prepath, "temp")
		filepath :=  filepath.Join("../../FRONTEND/frontend-service/static/temp/id-" + strconv.Itoa(int(userID)), files[i].Filename )
		out, err := os.Create(filepath)
		fmt.Println("final path is" + finalpath)

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

	var story data.Story
	story.PostedBy = dto.Username
	story.Media = medias[0]


	err = p.Service.CreateStory(&story)

	if err != nil {
		http.Error(
			rw,
			fmt.Sprintf("Error creating STORY %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	rw.WriteHeader(http.StatusOK)


}

func (s *StoryHandler) GetStories(rw http.ResponseWriter, r *http.Request) {
	s.L.Println("Handle GET Request for Posts")

	ls := data.GetStories()

	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal stories json", http.StatusInternalServerError)
	}
}

func (s *StoryHandler) GetStoriesForCurrentUser(rw http.ResponseWriter, r *http.Request) {
	// send whoami to auth service
	resp, err := UserCheck(r.Header.Get("Authorization"))
	if err != nil {
		s.L.Fatalln("There has been an error sending the /whoami request")
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
	stories := s.Service.GetAllStoriesForUser(dto.Username)
	err = stories.ToJSON(rw)
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

func (h *StoryHandler) GetAllStoriesForFeedForCurrentUser(rw http.ResponseWriter, r *http.Request) {
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

	allStories := h.Service.GetAllStoriesForFeed(allUsernames)

	fmt.Println(allStories)

	allStoriesJSON, _ := json.Marshal(allStories)

	_, err = rw.Write(allStoriesJSON)

	if err!=nil{
		fmt.Println("Error with Write!")
		return
	}

}
