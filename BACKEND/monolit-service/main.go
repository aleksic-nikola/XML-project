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
	"xml/monolit-service/constants"
	"xml/monolit-service/data"
	"xml/monolit-service/handlers"
	"xml/monolit-service/repository"
	"xml/monolit-service/service"
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
		fmt.Println("Successfully connected to profile-database")
	}

	return database
}

func initUserRepo(database *gorm.DB) *repository.UserRepository {
	return &repository.UserRepository{Database: database}
}

func initProductRepo(database *gorm.DB) *repository.ProductRepository {
	return &repository.ProductRepository{Database: database}
}

func initUserServices(repo *repository.UserRepository) *service.UserService {
	return &service.UserService{Repo: repo}
}

func initProductServices(repo *repository.ProductRepository) *service.ProductService {
	return &service.ProductService{Repo: repo}
}

func initUserHandler(user_service *service.UserService) *handlers.UserHandler {
	l := log.New(os.Stdout, "monolit-service", log.LstdFlags)
	return &handlers.UserHandler{L: l, Service: user_service}
}

func initProductHandler(product_service *service.ProductService) *handlers.ProductHandler {
	l := log.New(os.Stdout, "monolit-service", log.LstdFlags)
	return &handlers.ProductHandler{L: l, Service: product_service}
}

func main() {

	fmt.Println("hello monolit-service")

	database := initDB()
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	database.AutoMigrate(&data.User{})
	database.AutoMigrate(&data.Product{})
	database.AutoMigrate(&data.Image{})
	//database.Model(&data.Product{}).Related(&data.Image{}, "ProductRefer")

	user_repo := initUserRepo(database)
	user_service := initUserServices(user_repo)
	product_repo := initProductRepo(database)
	product_service := initProductServices(product_repo)

	uh := initUserHandler(user_service)
	ph := initProductHandler(product_service)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/getusers", uh.GetUsers)
	getRouter.HandleFunc("/getproducts", ph.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/users/add", uh.CreateUser)
	postRouter.HandleFunc("/products/add", ph.CreateProduct)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}))

	l := log.New(os.Stdout, "monolit-service ", log.LstdFlags)

	s := http.Server {
		Addr: constants.PORT,
		Handler : ch(sm),
		ErrorLog: l,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5* time.Second,
		IdleTimeout:  120* time.Second,
	}

	go func() {

		l.Println("Starting on server port 2902")
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
