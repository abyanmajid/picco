package main

import (
	"context"
	"log"
	"os"

	"github.com/abyanmajid/codemore.io/services/course/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := os.Getenv("DB_URI")
	if uri == "" {
		log.Fatal("Set your 'DB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	params := handlers.CourseServiceParameters{
		Port:       "50001",
		App:        "Course",
		Client:     client,
		Database:   "db",
		Collection: "courses",
	}

	handlers.ListenAndServeCourse(params)
}
