package constants

import (
	"github.com/joho/godotenv"
	"os"
)

var HOST = GetVariable("HOST")
var DOMAIN = GetVariable("DOMAIN")
var PORT = GetVariable("PORT")
var AUTH_SERVICE_URL = GetVariable("AUTH_SERVICE")
var CAMPAIGN_SERVICE_URL = GetVariable("CAMPAIGN_SERVICE")
var CONTENT_SERVICE_URL = GetVariable("CONTENT_SERVICE")
var INTERACTION_SERVICE_URL = GetVariable("INTERACTION_SERVICE")
var MONOLIT_SERVICE_URL = GetVariable("MONOLIT_SERVICE")
var PROFILE_SERVICE_URL = GetVariable("PROFILE_SERVICE")
var REQUEST_SERVICE_URL = GetVariable("REQUEST_SERVICE")
var SEARCH_SERVICE_URL = GetVariable("SEARCH_SERVICE")

func GetVariable(name string) string {

	godotenv.Load()
	if os.Getenv("DOCKERIZED") != "yes" {
		// running locally
		if name == "HOST" {
			return "localhost"
		}
		if name == "PORT" {
			return ":" + os.Getenv("LOCAL_PORT")
		}

		if name == "DOMAIN" {
			return "localhost"
		}
		if name == "AUTH_SERVICE" {
			return os.Getenv("AUTH_SERVICE_URL")
		}
		if name == "CAMPAIGN_SERVICE" {
			return os.Getenv("CAMPAIGN_SERVICE_URL")
		}
		if name == "CONTENT_SERVICE" {
			return os.Getenv("CONTENT_SERVICE_URL")
		}
		if name == "INTERACTION_SERVICE" {
			return os.Getenv("INTERACTION_SERVICE_URL")
		}
		if name == "MONOLIT_SERVICE" {
			return os.Getenv("MONOLIT_SERVICE_URL")
		}
		if name == "PROFILE_SERVICE" {
			return os.Getenv("PROFILE_SERVICE_URL")
		}
		if name == "REQUEST_SERVICE" {
			return os.Getenv("REQUEST_SERVICE_URL")
		}
		if name == "SEARCH_SERVICE" {
			return os.Getenv("SEARCH_SERVICE_URL")
		}


		return ""
	}

	// running in docker container
	if name == "HOST" {
		return os.Getenv("HOST_NAME")
	}
	if name == "PORT" {
		return ":" + os.Getenv("REQUEST_SERVICE_PORT")
	}

	if name == "DOMAIN" {
		return os.Getenv("REQUEST_SERVICE_DOMAIN")
	}

	if name == "AUTH_SERVICE" {
		return "auth-service:" + os.Getenv("AUTH_SERVICE_PORT")
	}
	if name == "CAMPAIGN_SERVICE" {
		return "campaign-service:" + os.Getenv("CAMPAIGN_SERVICE_PORT")
	}
	if name == "CONTENT_SERVICE" {
		return "content-service:" + os.Getenv("CONTENT_SERVICE_PORT")
	}
	if name == "INTERACTION_SERVICE" {
		return "interaction-service:" + os.Getenv("INTERACTION_SERVICE_PORT")
	}
	if name == "MONOLIT_SERVICE" {
		return "monolit-service:" + os.Getenv("MONOLIT_SERVICE_PORT")
	}
	if name == "PROFILE_SERVICE" {
		return "profile-service:" + os.Getenv("PROFILE_SERVICE_PORT")
	}
	if name == "REQUEST_SERVICE" {
		return "request-service:" + os.Getenv("REQUEST_SERVICE_PORT")
	}
	if name == "SEARCH_SERVICE" {
		return "search-service:" + os.Getenv("SEARCH_SERVICE_PORT")
	}

	return ""
}




