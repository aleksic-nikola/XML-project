package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
	"xml/content-service/handlers"
)

func main() {

	l := log.New(os.Stdout, "content-service ", log.LstdFlags)

	// post handler
	ph := handlers.NewPosts(l)
	sh := handlers.NewStories(l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/posts", ph.GetPosts)
	getRouter.HandleFunc("/stories", sh.GetStories)

	s := http.Server{
		Addr:         ":8080",      // configure the bind address
		Handler:      sm,                // set the default handler
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

	runtime.Goexit()
}