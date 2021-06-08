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
	"xml/interaction-service/handlers"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("hello")

	l := log.New(os.Stdout, "interaction-service", log.LstdFlags)

	msgH := handlers.NewMessages(l)
	msgwcH := handlers.NewMessagesWithContent(l)
	msgwotcH := handlers.NewMessagesWithOneTimeContent(l)

	profnotifH := handlers.NewProfileNotifications(l)
	postnotifH := handlers.NewPostNotifications(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/getmsgs", msgH.GetMessages).Methods("GET")
	getRouter.HandleFunc("/getmsgswithcontent", msgwcH.GetMessagesWithContent).Methods("GET")
	getRouter.HandleFunc("/getmsgswithonetimecontent", msgwotcH.GetMessagesWithOneTimeContent).Methods("GET")
	getRouter.HandleFunc("/getprofnotif", profnotifH.GetProfileNotifications).Methods("GET")
	getRouter.HandleFunc("/getpostnotif", postnotifH.GetPostNotifications).Methods("GET")

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}))

	s := http.Server{
		Addr:         ":5050",           // configure the bind address
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
