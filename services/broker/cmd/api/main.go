package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	utils "github.com/abyanmajid/codemore.io/services/broker/utils"
)

const PORT = "80"
const APP_NAME = "Broker"

func SetMicroservices(api *Service) {
	environment := os.Getenv("ENVIRONMENT")

	switch environment {
	case "development":
		api.UserEndpoint = "user:50001"
		api.CompilerEndpoint = "compiler:50001"
		api.JudgeEndpoint = "judge:50001"
		api.ContentFetcherEndpoint = "content-fetcher:50001"
		api.CourseEndpoint = "course:50001"
	case "production":
		api.UserEndpoint = os.Getenv("PRODUCTION_USER_ENDPOINT")
		api.CompilerEndpoint = os.Getenv("PRODUCTION_COMPILER_ENDPOINT")
		api.JudgeEndpoint = os.Getenv("PRODUCTION_JUDGE_ENDPOINT")
		api.ContentFetcherEndpoint = os.Getenv("PRODUCTION_CONTENT_ENDPOINT")
		api.CourseEndpoint = os.Getenv("PRODUCTION_COURSE_ENDPOINT")
	default:
		log.Fatal("The ENVIRONMENT environment variable is either not set or is not 'development' or 'production'")
	}
}

func main() {
	api := Service{
		Log: slog.New(utils.StructuredLogHandler(os.Stdout, APP_NAME)),
	}

	SetMicroservices(&api)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: api.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
