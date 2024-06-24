package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (api *Service) routes() http.Handler {
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
	router.Put("/user/id/{id}", api.HandleUpdateUserById)
	router.Delete("/user/id/{id}", api.HandleDeleteUserById)

	// Proxying routes for compiler service
	router.Post("/compiler/python", api.HandleCompilePython)
	router.Post("/compiler/java", api.HandleCompileJava)
	router.Post("/compiler/cpp", api.HandleCompileCpp)
	router.Post("/compiler/javascript", api.HandleCompileJavaScript)

	// Proxying routes for judge service
	router.Post("/judge/testcases/{task_id}", api.HandleCreateTestCase)
	router.Get("/judge/testcases/{task_id}", api.HandleGetAllTestCases)
	router.Get("/judge/testcases/{task_id}/{test_case_id}", api.HandleGetTestCase)
	router.Put("/judge/testcases/{task_id}/{test_case_id}", api.HandleUpdateTestCase)
	router.Delete("/judge/testcases/{task_id}/{test_case_id}", api.HandleDeleteTestCase)
	router.Delete("/judge/testcases/{task_id}", api.HandleDeleteAllTestCases)
	router.Post("/judge/tests/{task_id}", api.HandleRunTests)

	return router
}
