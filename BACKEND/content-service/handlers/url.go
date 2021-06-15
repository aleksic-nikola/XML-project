package handlers

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)


func GetVariable(service_name string) string {

	godotenv.Load()
	var retString string

	if os.Getenv("DOCKERIZED") != "yes" {
		retString = os.Getenv(strings.ToUpper(service_name) + "_SERVICE_URL")	
	} else {
		retString = "API_VARIABLE_HERE"
	}
	return retString	
}