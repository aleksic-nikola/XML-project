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
	"xml/campaign-service/constants"
	"xml/campaign-service/data"
	"xml/campaign-service/handlers"
	repository "xml/campaign-service/repository"
	"xml/campaign-service/service"
)

func initDB() *gorm.DB {

	godotenv.Load()
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	name := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// db connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", constants.HOST, user, name, password, dbport)
	println(dbURI + "*****")
	// open connection to db
	database, errx := gorm.Open(postgres.Open(dbURI))

	if errx != nil {
		log.Fatal(errx)
	} else {
		fmt.Println("Successfully connected to campaign-database")
	}

	return database
}

func initMonitoringRepo(database *gorm.DB) *repository.MonitoringReportRepository {
	return &repository.MonitoringReportRepository{Database: database}
}

func initMultiCampaignRepo(database *gorm.DB) *repository.MultiCampaignRepository {
	return &repository.MultiCampaignRepository{Database: database}
}

func initOneTimeCampaignRepo(database *gorm.DB) *repository.OneTimeCampaignRepository {
	return &repository.OneTimeCampaignRepository{Database: database}
}

func initOneTimeCampaignServices(repo *repository.OneTimeCampaignRepository) *service.OneTimeCampaignService {
	return &service.OneTimeCampaignService{Repo: repo}
}

func initMultiCampaignServices(repo *repository.MultiCampaignRepository) *service.MultiCampaignService {
	return &service.MultiCampaignService{Repo: repo}
}

func initMonitoringReportServices(repo *repository.MonitoringReportRepository) *service.MonitoringReportService {
	return &service.MonitoringReportService{Repo: repo}
}

func initMonitoringReportHandler(monitoringReport_service *service.MonitoringReportService) *handlers.MonitoringReportHandler {
	l := log.New(os.Stdout, "campaign-service", log.LstdFlags)
	return &handlers.MonitoringReportHandler{L: l, Service: monitoringReport_service}
}

func initMultiCampaignHandler(multiCampaign_service *service.MultiCampaignService) *handlers.MultiCampaignHandler {
	l := log.New(os.Stdout, "campaign-service", log.LstdFlags)
	return &handlers.MultiCampaignHandler{L: l, Service: multiCampaign_service}
}

func initOneTimeCampaignHandler(oneTimeCampaign_service * service.OneTimeCampaignService) *handlers.OneTimeCampaignHandler {
	l := log.New(os.Stdout, "campaign-service", log.LstdFlags)
	return &handlers.OneTimeCampaignHandler{L: l, Service: oneTimeCampaign_service}
}


func main() {

	fmt.Println("hello campaign-service")

	database := initDB()
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	database.AutoMigrate(&data.OneTimeCampaign{})
	database.AutoMigrate(&data.MultiCampaign{})
	database.AutoMigrate(&data.Campaign{})
	database.AutoMigrate(&data.MonitoringReport{})

	oneTimeCampaign_repo := initOneTimeCampaignRepo(database)
	oneTimeCampaign_service := initOneTimeCampaignServices(oneTimeCampaign_repo)

	multiCampaign_repo := initMultiCampaignRepo(database)
	multiCampaign_service := initMultiCampaignServices(multiCampaign_repo)

	monitoringReport_repo := initMonitoringRepo(database)
	monitoringReport_service := initMonitoringReportServices(monitoringReport_repo)

	otch := initOneTimeCampaignHandler(oneTimeCampaign_service)
	mch := initMultiCampaignHandler(multiCampaign_service)
	mrh := initMonitoringReportHandler(monitoringReport_service)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/monitoringreports", mrh.GetMonitoringReports)
	getRouter.HandleFunc("/multicampaigns", mch.GetMultiCampaigns)
	getRouter.HandleFunc("/onetimecampaigns", otch.GetOneTimeCampaigns)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/monitoringreports/add", mrh.CreateMonitoringReport)
	postRouter.HandleFunc("/multicampaings/add", mch.CreateMultiCampaign)
	postRouter.HandleFunc("/onetimecampaigns/add", otch.CreateOneTimeCampaign)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}))

	l := log.New(os.Stdout, "content-service ", log.LstdFlags)

	s := http.Server{
		Addr:         constants.PORT,      // configure the bind address
		Handler:      ch(sm),                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Starting server on port 7071")

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
