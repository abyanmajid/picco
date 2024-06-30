package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abyanmajid/codemore.io/services/course/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PORT = "80"
const APP_NAME = "Course"

type Service struct {
	Client *mongo.Client
}

func (api *Service) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	courseService := handlers.CourseService{Collection: api.Client.Database("db").Collection("courses")}

	// Course handlers
	router.Post("/course", courseService.HandleCreateCourse)
	router.Get("/course", courseService.HandleGetAllCourses)
	router.Get("/course/{title}", courseService.HandleGetCourseByTitle)
	router.Put("/course/{title}", courseService.HandleUpdateCourseByTitle)
	router.Delete("/course/{title}", courseService.HandleDeleteCourseByTitle)

	// Progression handlers

	return router
}

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

	api := Service{Client: client}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: api.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
