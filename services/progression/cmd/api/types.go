package main

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Mongo *mongo.Client
	Log   *slog.Logger
}
