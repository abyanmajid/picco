package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abyanmajid/codemore.io/services/broker/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

const PORT = "80"
const APP_NAME = "Broker"

type BrokerService struct {
	UserEndpoint           string
	CompilerEndpoint       string
	JudgeEndpoint          string
	ContentFetcherEndpoint string
	CourseEndpoint         string
}

func (api *BrokerService) setEndpoints() {
	api.UserEndpoint = "user:50001"
	api.CompilerEndpoint = "compiler:50001"
	api.JudgeEndpoint = "judge:50001"
	api.ContentFetcherEndpoint = "content-fetcher:50001"
	api.CourseEndpoint = "course:50001"
}

func (api *BrokerService) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Intialize services
	compilerService := handlers.CompilerService{Endpoint: api.CompilerEndpoint}
	judgeService := handlers.JudgeService{Endpoint: api.JudgeEndpoint}
	courseService := handlers.CourseService{Endpoint: api.CourseEndpoint}
	contentFetcherService := handlers.ContentFetcherService{Endpoint: api.ContentFetcherEndpoint}

	// Proxying routes for compiler service
	router.Post("/compiler/python", compilerService.HandleCompilePython)
	router.Post("/compiler/java", compilerService.HandleCompileJava)
	router.Post("/compiler/cpp", compilerService.HandleCompileCpp)
	router.Post("/compiler/javascript", compilerService.HandleCompileJavaScript)

	// Proxying routes for judge service
	router.Post("/judge", judgeService.HandleRunTests)

	// Proxying routes for course service
	router.Post("/course", courseService.HandleCreateCourse)
	router.Get("/course", courseService.HandleGetAllCourses)
	router.Get("/course/{title}", courseService.HandleGetCourseByTitle)
	router.Put("/course/{title}", courseService.HandleUpdateCourseByTitle)
	router.Delete("/course/{title}", courseService.HandleDeleteCourseByTitle)

	// Proxying routes for content fetcher service
	router.Post("/content-fetcher", contentFetcherService.HandleGetContent)

	return router
}

func main() {
	api := BrokerService{}

	api.setEndpoints()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: api.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
