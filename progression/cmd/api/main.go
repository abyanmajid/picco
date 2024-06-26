package main

import (
	"context"
	"log"
	"os"

	"github.com/abyanmajid/codemore.io/progression/utils"
)

const PORT = "50001"
const APP_NAME = "Judge"
const DEFAULT_DEVELOPMENT_DB_URL = "mongodb://mongo:27017"

func main() {
	environment := os.Getenv("ENVIRONMENT")
	var dbURL string
	var username string
	var password string

	switch environment {
	case "development":
		dbURL = DEFAULT_DEVELOPMENT_DB_URL
		username = "mongo"
		password = "mongo"
	case "production":
		dbURL = os.Getenv("PRODUCTION_DB_URL")
		username = os.Getenv("PRODUCTION_DB_USERNAME")
		password = os.Getenv("PRODUCTION_DB_PASSWORD")
	default:
		log.Fatal("The ENVIRONMENT environment variable is either not set or is not 'development' or 'production'")
	}

	ctx := context.TODO()
	mongoClient, err := utils.ConnectToDB(ctx, dbURL, username, password)
	if err != nil {
		return
	}
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			log.Panic()
		}
	}()

	ListenAndServe(mongoClient)
}