package main

import (
	"log/slog"

	"github.com/abyanmajid/codemore.io/services/course/proto/course"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	course.UnimplementedCourseServiceServer
	Mongo *mongo.Client
	Log   *slog.Logger
}

type Course struct {
	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	Creator     string   `bson:"creator"`
	Likes       int32    `bson:"likes"`
	Students    []string `bson:"students"`
	Topics      []string `bson:"topics"`
	Modules     []Module `bson:"modules"`
	UpdatedAt   string   `bson:"updated_at"`
	CreatedAt   string   `bson:"created_at"`
}

type Module struct {
	Id    int32  `bson:"id"`
	Title string `bson:"title"`
	Tasks []Task `bson:"tasks"`
}

type Task struct {
	Id   int32  `bson:"id"`
	Task string `bson:"task"`
	Type string `bson:"type"`
	Xp   int32  `bson:"xp"`
}
