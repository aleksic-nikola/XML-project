package main

import (
	"context"
	"fmt"
	gohandlers "github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
	"xml/profile-service/constants"
	"xml/profile-service/data"
	"xml/profile-service/handlers"
	"xml/profile-service/repository"
	"xml/profile-service/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	constants.PrintVars()
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

func initProfileRepo(database *gorm.DB) *repository.ProfileRepository {
	return &repository.ProfileRepository{Database: database}
}

func initAgentRepo(database *gorm.DB) *repository.AgentRepository {
	return &repository.AgentRepository{Database: database}
}

func initVerifiedRepo(database *gorm.DB) *repository.VerifiedRepository {
	return &repository.VerifiedRepository{Database: database}
}

func initProfileServices(repo *repository.ProfileRepository) *service.ProfileService {
	return &service.ProfileService{Repo: repo}
}

func initAgentServices(repo *repository.AgentRepository) *service.AgentService {
	return &service.AgentService{Repo: repo}
}

func initVerifiedServices(repo *repository.VerifiedRepository) *service.VerifiedService {
	return &service.VerifiedService{Repo: repo}
}

func initProfileHandler(service *service.ProfileService) *handlers.ProfileHandler {
	l := log.New(os.Stdout, "profile-service ", log.LstdFlags)
	return &handlers.ProfileHandler{L: l, Service: service}
}

func initAgentHandler(service *service.AgentService) *handlers.AgentHandler {
	l := log.New(os.Stdout, "profile-service ", log.LstdFlags)
	return &handlers.AgentHandler{L: l, Service: service}
}

func initVerifiedHandler(service *service.VerifiedService) *handlers.VerifiedHandler {
	l := log.New(os.Stdout, "profile-service ", log.LstdFlags)
	return &handlers.VerifiedHandler{L: l, Service: service}
}


func main() {

	fmt.Println("Hello profile-service")

	database := initDB()
	// defered closing
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	// make migrations to the db if they're not already created
	database.AutoMigrate(&data.Profile{})

	database.AutoMigrate(&data.Agent{})
	database.AutoMigrate(&data.Verified{})

	profile_repo := initProfileRepo(database)
	profile_service := initProfileServices(profile_repo)
	//agent_service := initAgentServices(repo)
	//verified_service := initVerifiedServices(repo)

	l := log.New(os.Stdout, "profile-service", log.LstdFlags)
	ph := initProfileHandler(profile_service)
	//agh := initAgentHandler(agent_service)
	//vh := initVerifiedHandler(verified_service)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/getprofiles", ph.GetProfiles)
	//getRouter.HandleFunc("/getagents", agh.GetAgents)
	//getRouter.HandleFunc("/getadmins", adh.GetAdmins)
	//getRouter.HandleFunc("/getverifieds", vh.GetVerifieds)
	getRouter.HandleFunc("/isuserpublic/{username}", ph.IsUserPublic)


	getRouter.HandleFunc("/getdata", ph.GetCurrent)

	getRouter.HandleFunc("/getAllFollowing", ph.GetAllFollowingUsernameBy)


	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/profile/create", ph.CreateProfile)
	postRouter.HandleFunc("/addprofile", ph.CreateProfile)
	postRouter.HandleFunc("/follow", ph.FollowAccount)
	postRouter.HandleFunc("/acceptFollow", ph.AcceptFollow)
	postRouter.HandleFunc("/editprofile", ph.EditProfileData)
	postRouter.HandleFunc("/editprivacysettings", ph.EditProfilePrivacySettings)
	postRouter.HandleFunc("/editnotifsettings", ph.EditProfileNotificationSettings)

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}))

	s := http.Server {
		Addr: constants.PORT,
		Handler : ch(sm),
		ErrorLog: l,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5* time.Second,
		IdleTimeout:  120* time.Second,
	}

	go func() {

		l.Println("Starting on server port 3030")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

	runtime.Goexit()


}
