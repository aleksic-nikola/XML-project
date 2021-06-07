package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
	"xml/request-service/handlers"
)

func main() {

	//env.Parse()

	fmt.Println("Hello, world.")

	l := log.New(os.Stdout, "auth-service ", log.LstdFlags )

	rh := handlers.NewRequest(l)
	mrh := handlers.NewMessageRequest(l)
	sch := handlers.NewSensitiveContentReportRequest(l)
	arrh := handlers.NewAgentRegistrationRequest(l)
	frh := handlers.NewFollowRequest(l)
	irh := handlers.NewInfluenceRequest(l)
	mrrh := handlers.NewMonitorReportRequest(l);
	vrh := handlers.NewVerificationRequest(l);

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/req", rh.GetRequests)
	getRouter.HandleFunc("/messReqs", mrh.GetMessageRequests)
	getRouter.HandleFunc("/sensitiveContent", sch.GetSensitiveContentReportRequests)
	getRouter.HandleFunc("/agentRegReqs", arrh.GetAgentRegistrationRequests)
	getRouter.HandleFunc("/followReqs", frh.GetFollowRequests)
	getRouter.HandleFunc("/influenceReqs", irh.GetInfluenceRequests)
	getRouter.HandleFunc("/monitorReportReqs", mrrh.GetMonitorReportRequests)
	getRouter.HandleFunc("/verificationReqs", vrh.GetVerificationRequests)


	s := http.Server {
		Addr: ":9090", // configure the bind address
		Handler: sm,  // set the default handler
		ErrorLog: l, // set the logger for the server
		ReadTimeout: 5 * time.Second, // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
		IdleTimeout: 120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Starting server on port 9090")
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