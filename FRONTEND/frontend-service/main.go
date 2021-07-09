package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))

	var port string

	if os.Getenv("DOCKERIZED") != "yes" {
		// local
		port = ":8000"
	} else {
		port = ":8000"
	}

	server := &http.Server{
		Addr:           port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Frontend running on port " + port)
	log.Fatal(server.ListenAndServe())
}
