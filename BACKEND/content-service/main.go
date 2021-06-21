package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
	"xml/content-service/constants"
	"xml/content-service/data"
	"xml/content-service/handlers"
	"xml/content-service/repository"
	"xml/content-service/service"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func printVariables() {
	fmt.Println("------------------------------")
	fmt.Println(handlers.GetVariable("auth"))
	fmt.Println(handlers.GetVariable("campaign"))
	fmt.Println(handlers.GetVariable("content"))
	fmt.Println(handlers.GetVariable("interaction"))
	fmt.Println(handlers.GetVariable("monolit"))
	fmt.Println(handlers.GetVariable("profile"))
	fmt.Println(handlers.GetVariable("request"))
	fmt.Println(handlers.GetVariable("search"))
	fmt.Println("------------------------------")

}

func initDB() *gorm.DB {

	godotenv.Load()
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	name := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// db connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", constants.HOST, user, name, password, dbport)

	// open connection to db
	database, errx := gorm.Open(postgres.Open(dbURI))

	if errx != nil {
		log.Fatal(errx)
	} else {
		fmt.Println("Successfully connected to content-database")
	}

	return database
}

// INIT REPOS
func initPostRepo(database *gorm.DB) *repository.PostRepository {
	return &repository.PostRepository{Database: database}
}

func initStoryRepo(database *gorm.DB) *repository.StoryRepository {
	return &repository.StoryRepository{Database: database}
}

func initMediaRepo(database *gorm.DB) *repository.MediaRepository {
	return &repository.MediaRepository{Database: database}
}

func initCommentRepo(database *gorm.DB) *repository.CommentRepository {
	return &repository.CommentRepository{Database: database}
}

func initLocationRepo(database *gorm.DB) *repository.LocationRepository {
	return &repository.LocationRepository{Database: database}
}

// INIT SERVICES

func initPostServices(repo *repository.PostRepository) *service.PostService {
	return &service.PostService{Repo: repo}
}

func initStoryServices(repo *repository.StoryRepository) *service.StoryService {
	return &service.StoryService{Repo: repo}
}

func initMediaServices(repo *repository.MediaRepository) *service.MediaService {
	return &service.MediaService{Repo: repo}
}

func initCommentServices(repo *repository.CommentRepository) *service.CommentService {
	return &service.CommentService{Repo: repo}
}

func initLocationServices(repo *repository.LocationRepository) *service.LocationService {
	return &service.LocationService{Repo: repo}
}

// INIT HANDLERS

func initPostHandler(service *service.PostService) *handlers.PostHandler {
	l := log.New(os.Stdout, "content-service ", log.LstdFlags)
	return &handlers.PostHandler{L: l, Service: service}
}

func initStoryHandler(service *service.StoryService) *handlers.StoryHandler {
	l := log.New(os.Stdout, "content-service ", log.LstdFlags)
	return &handlers.StoryHandler{L: l, Service: service}
}

func initMediaHandler(service *service.MediaService) *handlers.MediaHandler {
	l := log.New(os.Stdout, "content-service ", log.LstdFlags)
	return &handlers.MediaHandler{L: l, Service: service}
}

func initCommentHandler(service *service.CommentService) *handlers.CommentHandler {
	l := log.New(os.Stdout, "content-service ", log.LstdFlags)
	return &handlers.CommentHandler{L: l, Service: service}
}

func initLocationHandler(service *service.LocationService) *handlers.LocationHandler {
	l := log.New(os.Stdout, "content-service ", log.LstdFlags)
	return &handlers.LocationHandler{L: l, Service: service}
}

func main() {

	printVariables()

	l := log.New(os.Stdout, "content-service ", log.LstdFlags)

	database := initDB()
	// defered closing
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	// make migrations to the db if they're not already created
	database.AutoMigrate(&data.Post{})
	database.AutoMigrate(&data.Story{})
	database.AutoMigrate(&data.Media{})
	database.AutoMigrate(&data.Comment{})

	// INCLUDE COMMENT, MEDIA, LOCATION repo, service, handler if needed
	post_repo := initPostRepo(database)
	story_repo := initStoryRepo(database)

	post_service := initPostServices(post_repo)
	story_service := initStoryServices(story_repo)

	// post handler
	ph := initPostHandler(post_service)
	sh := initStoryHandler(story_service)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/posts", ph.GetPosts)
	getRouter.HandleFunc("/stories", sh.GetStories)
	getRouter.HandleFunc("/getpostsbyuser/{username}", ph.GetPostsByUser)
	getRouter.HandleFunc("/current/posts", ph.GetPostsForCurrentUser)
	getRouter.HandleFunc("/current/stories", sh.GetStoriesForCurrentUser)
	getRouter.HandleFunc("/current/likedposts", ph.GetLikedPostsByUser)
	getRouter.HandleFunc("/current/dislikedposts", ph.GetDislikedPostsByUser)
	getRouter.HandleFunc("/current/getAllPostsForFeed", ph.GetAllPostsForFeedForCurrentUser)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/post/add", ph.CreatePost)
	postRouter.HandleFunc("/post/upload", ph.UploadPost)
	postRouter.HandleFunc("/story/add", sh.CreateStory)

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}))

	s := http.Server{
		Addr:         constants.GetVariable("PORT"),           // configure the bind address
		Handler:      ch(sm),            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

	runtime.Goexit()
}
