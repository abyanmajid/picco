package main

import (
	"log"
	"os"

	"github.com/abyanmajid/codemore.io/user/utils"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const PORT = "50001"
const DEFAULT_DEVELOPMENT_DB_URL = "host=postgres port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"
const APP_NAME = "User"

func main() {
	environment := os.Getenv("ENVIRONMENT")
	var dbURL string

	switch environment {
	case "development":
		dbURL = DEFAULT_DEVELOPMENT_DB_URL
	case "production":
		dbURL = os.Getenv("PRODUCTION_DB_URL")
	default:
		log.Fatal("The ENVIRONMENT environment variable is either not set or is not 'development' or 'production'")
	}

	conn := utils.ConnectToDB(dbURL)
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	ListenAndServe(conn)
}
