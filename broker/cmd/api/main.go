package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT = "80"

func main() {
	environment := os.Getenv("ENVIRONMENT")
	var userEndpoint string

	switch environment {
	case "development":
		userEndpoint = "http://user"
	case "production":
		userEndpoint = os.Getenv("PRODUCTION_USER_ENDPOINT")
	default:
		log.Fatal("The ENVIRONMENT environment variable is either not set or is not 'development' or 'production'")
	}

	api := Config{
		UserEndpoint: userEndpoint,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: api.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
