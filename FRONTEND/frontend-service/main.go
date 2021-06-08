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
	mux.Handle("/", http.FileServer(http.Dir("FRONTEND/frontend-service/static")))

	
	server := &http.Server{
		Addr:           ":8000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}