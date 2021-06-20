package main

import (
	"context"
	"fmt"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
	"xml/auth-service/constants"
	"xml/auth-service/data"
	"xml/auth-service/handlers"
	"xml/auth-service/repository"
	"xml/auth-service/service"
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
	println(dbURI + "*****")
	// open connection to db
	database, errx := gorm.Open(postgres.Open(dbURI))

	if errx != nil {
		log.Fatal(errx)
	} else {
		fmt.Println("Successfully connected to auth-database")
	}

	return database
}

func initRepo(database *gorm.DB) *repository.UserRepository {
	return &repository.UserRepository{Database: database}
}

func initServices(repo *repository.UserRepository) *service.UserService {
	return &service.UserService{Repo: repo}
}

func initHandler(service *service.UserService) *handlers.UserHandler {
	l := log.New(os.Stdout, "auth-service ", log.LstdFlags)
	return &handlers.UserHandler{L: l, Service: service}
}

func main() {
	constants.PrintVars()
	database := initDB()
	// defered closing
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	// make migrations to the db if they're not already created
	database.AutoMigrate(&data.User{})

	repo := initRepo(database)
	service := initServices(repo)
	uh := initHandler(service)

	fmt.Println("hello")

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", uh.GetUsers)
	getRouter.HandleFunc("/whoami", uh.WhoAmI)
	getRouter.HandleFunc("/getuser/{username}", uh.GetUserByUsername)
	getRouter.Use(uh.AuthMiddleware)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/register", uh.CreateUser)
	postRouter.HandleFunc("/login", uh.Login)

	postRouter.HandleFunc("/edituser", uh.EditUserData)
	postRouter.HandleFunc("/changepw", uh.ChangePassowrd)

	getRouter.HandleFunc("/getdata", uh.GetCurrent)

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}))

	l := log.New(os.Stdout, "auth-service ", log.LstdFlags)

	s := http.Server{
		Addr:         constants.PORT, //":9090", //":8888",
		Handler:      ch(sm),
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {

		l.Println("Starting on server port 9090")
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
