package main

import (
	"log/slog"

	"github.com/abyanmajid/codemore.io/content/proto/content"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	content.UnimplementedContentServiceServer
	Mongo *mongo.Client
	Log   *slog.Logger
}

type Course struct {
	Id        string   `bson:"_id,omitempty"`
	Title     string   `bson:"title"`
	Creator   string   `bson:"creator"`
	Likes     int32    `bson:"likes"`
	Topics    []string `bson:"topics"`
	Modules   []Module `bson:"modules"`
	UpdatedAt string   `bson:"updated_at"`
	CreatedAt string   `bson:"created_at"`
}

type Module struct {
	Id    string `bson:"_id"`
	Title string `bson:"title"`
	Tasks []Task `bson:"tasks"`
}

type Task struct {
	Id    string `bson:"_id"`
	Title string `bson:"title"`
	Mdx   string `bson:"mdx"`
}
