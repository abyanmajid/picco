package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(ctx context.Context, dbURL string, username string, password string) (*mongo.Client, error) {
	log.Println("Starting database connection")

	// Set server API options
	opts := options.Client().ApplyURI(dbURL)
	opts.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	log.Println("Username:", username)
	log.Println("Password:", password)

	// Connect to the MongoDB server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v\n", err)
		return nil, err
	}
	log.Println("Connected to MongoDB server successfully")

	return client, nil
}
