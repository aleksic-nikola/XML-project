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
	"xml/search-service/constants"
	"xml/search-service/data"
	"xml/search-service/handlers"
	"xml/search-service/repository"
	"xml/search-service/service"
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
		fmt.Println("Successfully connected to search-database")
	}

	return database
}

func initQueryRepo(database *gorm.DB) *repository.QueryRepository  {
	return &repository.QueryRepository{Database: database}
}

func initQueryServices(repo *repository.QueryRepository) *service.QueryService {
	return &service.QueryService{Repo: repo}
}

func initQueryHandler(service *service.QueryService) *handlers.QueryHandler {
	l := log.New(os.Stdout, "search-service ", log.LstdFlags)
	return &handlers.QueryHandler{L: l, Service: service}
}

func main() {

	fmt.Println("hello search-service")

	database := initDB()
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	database.AutoMigrate(&data.Query{})

	queryRepo := initQueryRepo(database)
	queryService := initQueryServices(queryRepo)
	sh := initQueryHandler(queryService)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/getqueries", sh.GetQueries)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/search/addquery", sh.CreateQuery)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	gohandlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"})

	l := log.New(os.Stdout, "search-service", log.LstdFlags)

	s := http.Server {
		Addr: constants.PORT,
		Handler : ch(sm),
		ErrorLog: l,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5* time.Second,
		IdleTimeout:  120* time.Second,
	}

	go func() {

		l.Println("Starting on server port 9494")
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