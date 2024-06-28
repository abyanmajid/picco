package main

import (
	"log"
	"os"
)

const PORT = "50001"
const APP_NAME = "Judge"

func main() {
	environment := os.Getenv("ENVIRONMENT")
	var compilerEndpoint string
	var contentFetcherEndpoint string

	switch environment {
	case "development":
		compilerEndpoint = "compiler:50001"
		contentFetcherEndpoint = "content-fetcher:50001"
	case "production":
		compilerEndpoint = os.Getenv("PRODUCTION_COMPILER_ENDPOINT")
		contentFetcherEndpoint = os.Getenv("PRODUCTION_CONTENT_FETCHER_ENDPOINT")
	default:
		log.Fatal("The ENVIRONMENT environment variable is either not set or is not 'development' or 'production'")
	}

	ListenAndServe(compilerEndpoint, contentFetcherEndpoint)
}
