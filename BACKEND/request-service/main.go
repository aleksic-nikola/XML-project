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
	"xml/request-service/data"
	"xml/request-service/handlers"
	"xml/request-service/repository"
	"xml/request-service/service"
)

func initDB() *gorm.DB {

	godotenv.Load()
	host := os.Getenv("HOST")
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
		fmt.Println("Successfully connected to request-database")
	}

	return database
}

// Repository

func initFollowRequestRepository(database *gorm.DB) *repository.FollowRequestRepository {
	return &repository.FollowRequestRepository{Database: database}
}

func initAgentRegistrationRequestRepository(database *gorm.DB) *repository.AgentRegistrationRequestRepository{
	return &repository.AgentRegistrationRequestRepository{Database: database}
}

func initInfluenceRequestRepository(database *gorm.DB) *repository.InfluenceRequestRepository{
	return &repository.InfluenceRequestRepository{Database: database}
}

func initMessageRequestRepository(database *gorm.DB) *repository.MessageRequestRepository{
	return &repository.MessageRequestRepository{Database: database}
}

func initMonitorReportRequestRepository(database *gorm.DB) *repository.MonitorReportRequestRepository{
	return &repository.MonitorReportRequestRepository{Database: database}
}

func initSensitiveContentReportRequestRepository(database *gorm.DB) *repository.SensitiveContentReportRequestRepository{
	return &repository.SensitiveContentReportRequestRepository{Database: database}
}

func initVerificationRequestRepository(database *gorm.DB) *repository.VerificationRequestRepository{
	return &repository.VerificationRequestRepository{Database: database}
}
//-------------------------------------------------------



//Services

func initFollowRequestService(repo *repository.FollowRequestRepository) *service.FollowRequestService {
	return &service.FollowRequestService{Repo: repo}
}

func initAgentRegistrationRequestService(repo *repository.AgentRegistrationRequestRepository) *service.AgentRegistrationRequestService {
	return &service.AgentRegistrationRequestService{Repo: repo}
}

func initInfluenceRequestService(repo *repository.InfluenceRequestRepository) *service.InfluenceRequestService {
	return &service.InfluenceRequestService{Repo: repo}
}

func initMessageRequestService(repo *repository.MessageRequestRepository) *service.MessageRequestService {
	return &service.MessageRequestService{Repo: repo}
}

func initMonitorReportRequestService(repo *repository.MonitorReportRequestRepository) *service.MonitorReportRequestService {
	return &service.MonitorReportRequestService{Repo: repo}
}

func initSensitiveContentReportRequestService(repo *repository.SensitiveContentReportRequestRepository) *service.SensitiveContentReportRequestService {
	return &service.SensitiveContentReportRequestService{Repo: repo}
}

func initVerificationRequestService(repo *repository.VerificationRequestRepository) *service.VerificationRequestService {
	return &service.VerificationRequestService{Repo: repo}
}

//------------------------------------------------------------


//Handlers

func initFollowRequestHandler(service *service.FollowRequestService) *handlers.FollowRequestHandler {
	l := log.New(os.Stdout, "request-service ", log.LstdFlags)
	return &handlers.FollowRequestHandler{L: l, Service: service}
}

func initAgentRegistrationRequestHandler(service *service.AgentRegistrationRequestService) *handlers.AgentRegistrationRequestHandler {
	l := log.New(os.Stdout, "request-service ", log.LstdFlags)
	return &handlers.AgentRegistrationRequestHandler{L: l, Service: service}
}

func initInfluenceRequestHandler(service *service.InfluenceRequestService) *handlers.InfluenceRequestHandler {
	l := log.New(os.Stdout, "request-service ", log.LstdFlags)
	return &handlers.InfluenceRequestHandler{L: l, Service: service}
}

func initMessageRequestHandler(service *service.MessageRequestService) *handlers.MessageRequestHandler {
	l := log.New(os.Stdout, "request-service ", log.LstdFlags)
	return &handlers.MessageRequestHandler{L: l, Service: service}
}

func initMonitorReportRequestHandler(service *service.MonitorReportRequestService) *handlers.MonitorReportRequestHandler {
	l := log.New(os.Stdout, "request-service ", log.LstdFlags)
	return &handlers.MonitorReportRequestHandler{L: l, Service: service}
}

func initSensitiveContentReportRequestHandler(service *service.SensitiveContentReportRequestService) *handlers.SensitiveContentReportRequestHandler {
	l := log.New(os.Stdout, "request-service ", log.LstdFlags)
	return &handlers.SensitiveContentReportRequestHandler{L: l, Service: service}
}

func initVerificationRequestHandler(service *service.VerificationRequestService) *handlers.VerificationRequestHandler {
	l := log.New(os.Stdout, "request-service ", log.LstdFlags)
	return &handlers.VerificationRequestHandler{L: l, Service: service}
}



func authorized(h http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		jwtToken := r.Header.Get("Authorization")
		fmt.Println("BIOOO OVDEEE")
		resp, err := handlers.UserCheck(jwtToken)
		if err != nil {
			fmt.Println("BIOOO IIIIIIIIIIIII OVDEEE")

			//p.L.Fatalln("There has been an error sending the /whoami request")
			//http.Error(rw, "")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		if(resp.StatusCode != 200){ //biće 401 ako je greska, nije dobar token
			http.Error(rw, "Token error, unauthorized!", http.StatusUnauthorized)
			return
		}
		//code before
		h.ServeHTTP(rw,r)
		//code after
	})
}





func main() {

	//env.Parse()
	database := initDB()
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	database.AutoMigrate(&data.FollowRequest{})
	database.AutoMigrate(&data.AgentRegistrationRequest{})
	database.AutoMigrate(&data.MessageRequest{})
	database.AutoMigrate(&data.MonitorReportRequest{})
	database.AutoMigrate(&data.SensitiveContentReportRequest{})
	database.AutoMigrate(&data.VerificationRequest{})
	database.AutoMigrate(&data.InfluenceRequest{})


	//
	followRequestRepo := initFollowRequestRepository(database)
	followRequestService := initFollowRequestService(followRequestRepo)

	agentRegistrationRequestRepo := initAgentRegistrationRequestRepository(database)
	agentRegistrationRequestService := initAgentRegistrationRequestService(agentRegistrationRequestRepo)

	influenceRequestRepo := initInfluenceRequestRepository(database)
	influenceRequestService := initInfluenceRequestService(influenceRequestRepo)

	messageRequestRepo := initMessageRequestRepository(database)
	messageRequestService := initMessageRequestService(messageRequestRepo)

	monitorReportRequestRepo := initMonitorReportRequestRepository(database)
	monitorReportRequestService := initMonitorReportRequestService(monitorReportRequestRepo)

	sensitiveContentReportRequestRepo := initSensitiveContentReportRequestRepository(database)
	sensitiveContentReportRequestService := initSensitiveContentReportRequestService(sensitiveContentReportRequestRepo)

	verificationRequestRepo := initVerificationRequestRepository(database)
	verificationRequestService :=initVerificationRequestService(verificationRequestRepo)


	//handlers

	frh := initFollowRequestHandler(followRequestService)
	arrh := initAgentRegistrationRequestHandler(agentRegistrationRequestService)
	irh := initInfluenceRequestHandler(influenceRequestService)
	mrh := initMessageRequestHandler(messageRequestService)
	mrrh := initMonitorReportRequestHandler(monitorReportRequestService)
	scrrh := initSensitiveContentReportRequestHandler(sensitiveContentReportRequestService)
	vrh := initVerificationRequestHandler(verificationRequestService)



	fmt.Println("Hello, world.")


	//rh := handlers.NewRequest(l)
	//mrh := handlers.NewMessageRequest(l)
	//sch := handlers.NewSensitiveContentReportRequest(l)
	//arrh := handlers.NewAgentRegistrationRequest(l)
	//frh := handlers.NewFollowRequest(l)
	//irh := handlers.NewInfluenceRequest(l)
	//mrrh := handlers.NewMonitorReportRequest(l);
	//vrh := handlers.NewVerificationRequest(l);

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/messReqs", mrh.GetMessageRequests)
	getRouter.HandleFunc("/sensitiveContentReqs", scrrh.GetSensitiveContentReportRequests)
	getRouter.HandleFunc("/agentRegistrationReqs", arrh.GetAgentRegistrationRequests)
	getRouter.HandleFunc("/followReqs", frh.GetFollowRequests)
	getRouter.HandleFunc("/followReqs/getMy", authorized(frh.GetMyFollowRequests))

	getRouter.HandleFunc("/influenceReqs", irh.GetInfluenceRequests)
	getRouter.HandleFunc("/monitorReportReqs", mrrh.GetMonitorReportRequests)
	getRouter.HandleFunc("/verificationReqs", vrh.GetVerificationRequests)

	postRouter := sm.Methods(http.MethodPost).Subrouter()

	postRouter.HandleFunc("/followReqs/add", frh.CreateFollowRequest)
	postRouter.HandleFunc("/followReqs/accept", frh.AcceptFollowRequest)

	postRouter.HandleFunc("/agentRegistrationReqs/add", scrrh.CreateSensitiveContentReportRequest)
	postRouter.HandleFunc("/influenceReqs/add", irh.CreateInfluenceRequest)
	postRouter.HandleFunc("/messReqs/add", mrh.CreateMessageRequest)
	postRouter.HandleFunc("/monitorReportReqs/add", mrrh.CreateMonitorReportRequest)
	postRouter.HandleFunc("/sensitiveContentReqs/add", scrrh.CreateSensitiveContentReportRequest)
	postRouter.HandleFunc("/verificationReqs/add", vrh.CreateVerificationRequest)


	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}))


	l := log.New(os.Stdout, "request-service ", log.LstdFlags )

	s := http.Server {
		Addr: ":9211", // configure the bind address
		Handler: ch(sm),  // set the default handler
		ErrorLog: l, // set the logger for the server
		ReadTimeout: 5 * time.Second, // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
		IdleTimeout: 120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Starting server on port 9211")
		err :=  s.ListenAndServe()
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