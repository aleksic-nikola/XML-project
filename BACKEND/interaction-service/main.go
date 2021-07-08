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
	"xml/interaction-service/constants"
	"xml/interaction-service/data"
	"xml/interaction-service/handlers"
	"xml/interaction-service/repository"
	"xml/interaction-service/service"

	"gorm.io/gorm"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)

func initDB() *gorm.DB {

	godotenv.Load()
	host := constants.HOST
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	name := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// db connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, name, password, dbport)

	// open connection to db
	database, errx := gorm.Open(postgres.Open(dbURI))

	if errx != nil {
		log.Fatal(errx)
	} else {
		fmt.Println("Successfully connected to auth-database")
	}

	return database
}

//INIT REPOS
func initMessageRepo(database *gorm.DB) *repository.MessageRepository {
	return &repository.MessageRepository{Database: database}
}

func initMessageWithContentRepo(database *gorm.DB) *repository.MessageWithContentRepository {
	return &repository.MessageWithContentRepository{Database: database}
}

func initMessageWithOneTimeContentRepo(database *gorm.DB) *repository.MessageWithOneTimeContentRepository {
	return &repository.MessageWithOneTimeContentRepository{Database: database}
}

func initNotificationRepo(database *gorm.DB) *repository.NotificationRepository {
	return &repository.NotificationRepository{Database: database}
}

func initPostNotificationRepo(database *gorm.DB) *repository.PostNotificationRepository {
	return &repository.PostNotificationRepository{Database: database}
}

func initProfileNotificationRepo(database *gorm.DB) *repository.ProfileNotificationRepository {
	return &repository.ProfileNotificationRepository{Database: database}
}

//INIT SERVICES
func initMessageService(repo *repository.MessageRepository) *service.MessageService {
	return &service.MessageService{Repo: repo}
}

func initMessageWithContentService(repo *repository.MessageWithContentRepository) *service.MessageWithContentService {
	return &service.MessageWithContentService{Repo: repo}
}

func initMessageWithOneTimeContentService(repo *repository.MessageWithOneTimeContentRepository) *service.MessageWithOneTimeContentService {
	return &service.MessageWithOneTimeContentService{Repo: repo}
}

func initNotificationService(repo *repository.NotificationRepository) *service.NotificationService {
	return &service.NotificationService{Repo: repo}
}

func initPostNotificationService(repo *repository.PostNotificationRepository) *service.PostNotificationService {
	return &service.PostNotificationService{Repo: repo}
}

func initProfileNotificationService(repo *repository.ProfileNotificationRepository) *service.ProfileNotificationService {
	return &service.ProfileNotificationService{Repo: repo}
}

//HANDLERS
func initMessageHandler(service *service.MessageService) *handlers.MessageHandler {
	l := log.New(os.Stdout, "interaction-service ", log.LstdFlags)
	return &handlers.MessageHandler{L: l, Service: service}
}

func initMessageWithContentHandler(service *service.MessageWithContentService) *handlers.MessageWithContentHandler {
	l := log.New(os.Stdout, "interaction-service ", log.LstdFlags)
	return &handlers.MessageWithContentHandler{L: l, Service: service}
}

func initMessageWithOneTimeContentHandler(service *service.MessageWithOneTimeContentService) *handlers.MessageWithOneTimeContentHandler {
	l := log.New(os.Stdout, "interaction-service ", log.LstdFlags)
	return &handlers.MessageWithOneTimeContentHandler{L: l, Service: service}
}

func initNotificationHandler(service *service.NotificationService) *handlers.NotificationHandler {
	l := log.New(os.Stdout, "interaction-service ", log.LstdFlags)
	return &handlers.NotificationHandler{L: l, Service: service}
}

func initPostNotificationHandler(service *service.PostNotificationService) *handlers.PostNotificationHandler {
	l := log.New(os.Stdout, "interaction-service ", log.LstdFlags)
	return &handlers.PostNotificationHandler{L: l, Service: service}
}

func initProfileNotificationHandler(service *service.ProfileNotificationService) *handlers.ProfileNotificationHandler {
	l := log.New(os.Stdout, "interaction-service ", log.LstdFlags)
	return &handlers.ProfileNotificationHandler{L: l, Service: service}
}

func main() {

	fmt.Println("hello-interaction-service")

	database := initDB()

	// defered closing
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// make migrations to the db if they're not already created
	database.AutoMigrate(&data.Message{})
	database.AutoMigrate(&data.MessageWithContent{})
	database.AutoMigrate(&data.MessageWithOneTimeContent{})

	database.AutoMigrate(&data.Notification{})
	database.AutoMigrate(&data.PostNotification{})
	database.AutoMigrate(&data.ProfileNotification{})

	message_repo := initMessageRepo(database)
	message_service := initMessageService(message_repo)

	messagewithcontent_repo := initMessageWithContentRepo(database)
	messagewithcontent_service := initMessageWithContentService(messagewithcontent_repo)

	messagewithonetimecontent_repo := initMessageWithOneTimeContentRepo(database)
	messagewithonetimecontent_service := initMessageWithOneTimeContentService(messagewithonetimecontent_repo)

	notification_repo := initNotificationRepo(database)
	notification_service := initNotificationService(notification_repo)

	profnotification_repo := initProfileNotificationRepo(database)
	profnotification_service := initProfileNotificationService(profnotification_repo)

	postnotification_repo := initPostNotificationRepo(database)
	postnotification_service := initPostNotificationService(postnotification_repo)

	l := log.New(os.Stdout, "interaction-service", log.LstdFlags)

	msgH := initMessageHandler(message_service)
	msgwcH := initMessageWithContentHandler(messagewithcontent_service)
	msgwotcH := initMessageWithOneTimeContentHandler(messagewithonetimecontent_service)

	notifH := initNotificationHandler(notification_service)
	profnotifH := initProfileNotificationHandler(profnotification_service)
	postnotifH := initPostNotificationHandler(postnotification_service)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/getmsgs", msgH.GetMessages).Methods("GET")
	getRouter.HandleFunc("/getmsgswithcontent", msgwcH.GetMessagesWithContent).Methods("GET")
	getRouter.HandleFunc("/getmsgswithonetimecontent", msgwotcH.GetMessagesWithOneTimeContent).Methods("GET")
	getRouter.HandleFunc("/getnotif", notifH.GetNotifications).Methods("GET")
	getRouter.HandleFunc("/getprofnotif", profnotifH.GetProfileNotifications).Methods("GET")
	//getRouter.HandleFunc("/getpostnotif", postnotifH.GetPostNotifications).Methods("GET")
	getRouter.HandleFunc("/getunreadpostnotif", postnotifH.GetUnreadPostNotif).Methods("GET")

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/message/add", msgH.CreateMessage)
	postRouter.HandleFunc("/messagewithcontent/add", msgwcH.CreateMessageWithContent)
	postRouter.HandleFunc("/messagewithonetimecontent/add", msgwotcH.CreateMessageWithOneTimeContent)

	postRouter.HandleFunc("/postnotification/add", postnotifH.CreatePostNotification)
	postRouter.HandleFunc("/profilenotification/add", profnotifH.CreateProfileNotification)

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}))

	s := http.Server{
		Addr:         constants.PORT,           // configure the bind address
		Handler:      ch(sm),            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Starting server on port 5050")
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
