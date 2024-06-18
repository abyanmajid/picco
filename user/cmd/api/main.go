package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB *database.Queries
}

func main() {
	port, dbURL := loadEnv()

}

func loadEnv() (string, string) {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	return port, dbURL
}
