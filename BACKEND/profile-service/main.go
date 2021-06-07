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
	"xml/profile-serivce/handlers"
)

func main() {

	fmt.Println("Hello profile-service")

	l := log.New(os.Stdout, "profile-service", log.LstdFlags)
	ph := handlers.NewProfiles(l)
	agh := handlers.NewAgents(l)
	adh := handlers.NewAdmins(l)
	vh := handlers.NewVerifieds(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/getprofiles", ph.GetProfiles)
	getRouter.HandleFunc("/getagents", agh.GetAgents)
	getRouter.HandleFunc("/getadmins", adh.GetAdmins)
	getRouter.HandleFunc("/getverifieds", vh.GetVerifieds)

	s := http.Server {
		Addr: ":3030",
		Handler : sm,
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
