package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (api *Config) routes() http.Handler {
	router := chi.NewRouter()

	// specify who is allowed to connect
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Proxying routes for user service
	router.Post("/user", api.HandleCreateUser)
	router.Get("/user", api.HandleGetAllUsers)
	router.Get("/user/id/{id}", api.HandleGetUserById)
	router.Get("/user/email/{email}", api.HandleGetUserByEmail)
	router.Patch("/user/id/{id}", api.HandleUpdateUserById)
	router.Delete("/user/delete/{id}", api.HandleDeleteUserById)

	return router
}
