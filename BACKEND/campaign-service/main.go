package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
	"xml/campaign-service/handlers"
)

func main() {

	l := log.New(os.Stdout, "content-service ", log.LstdFlags)

	// post handler
	mrh := handlers.NewMonitoringReports(l)
	mch := handlers.NewMultiCampaigns(l)
	otch := handlers.NewOneTimeCampaigns(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/monitoringreports", mrh.GetMonitoringReports)
	getRouter.HandleFunc("/multicampaigns", mch.GetMultiCampaigns)
	getRouter.HandleFunc("/onetimecampaigns", otch.GetOneTimeCampaigns)

	s := http.Server{
		Addr:         ":7071",      // configure the bind address
		Handler:      sm,                // set the default handler
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
