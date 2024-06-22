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

	// Route for health check
	router.Get("/health", api.HandleHealth)

	// Routes for user-related operations
	router.Post("/users", api.HandleCreateUser)

	// Routes for auth-related operations
	router.Post("/auth/login", api.HandleLogin)
	// router.Get("/auth/refresh", api.RefreshToken)
	// router.Get("/auth/logout", api.HandleLogout)

	return router
}
