package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	utils "github.com/abyanmajid/codemore.io/broker/utils"
)

const PORT = "80"

func main() {
	environment := os.Getenv("ENVIRONMENT")

	api := Config{
		Log: slog.New(utils.StructuredLogHandler(os.Stdout)),
	}

	switch environment {
	case "development":
		api.UserEndpoint = "user:50001"
	case "production":
		api.UserEndpoint = os.Getenv("PRODUCTION_USER_ENDPOINT")
	default:
		log.Fatal("The ENVIRONMENT environment variable is either not set or is not 'development' or 'production'")
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
