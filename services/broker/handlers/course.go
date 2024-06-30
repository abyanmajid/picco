package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "log/slog"

	"github.com/abyanmajid/codemore.io/services/broker/proto/course"
	"github.com/abyanmajid/codemore.io/services/broker/utils"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CourseService struct {
	Endpoint string
}

type CourseServiceClient struct {
	Client course.CourseServiceClient
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

type Course struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	CreatorId   string   `json:"creator_id"`
	Likes       int32    `json:"likes"`
	Students    []string `json:"students"`
	Topics      []string `json:"topics"`
	Modules     []Module `json:"modules"`
	UpdatedAt   string   `json:"updated_at"`
	CreatedAt   string   `json:"created_at"`
}

type Module struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id   string `json:"id"`
	Task string `json:"task"`
	Type string `json:"type"`
	Xp   int32  `json:"xp"`
}

func (api *CourseService) getCourseServiceClient() (*CourseServiceClient, error) {

	conn, err := grpc.NewClient(api.Endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := course.NewCourseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	return &CourseServiceClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}

func (api *CourseService) HandleCreateCourse(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling CreateCourse request....")

	var requestPayload Course
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	client, err := api.getCourseServiceClient()
	if err != nil {
		log.Error("Failed to get gRPC client", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	log.Info("Creating new course", "title", requestPayload.Title)
	_, err = client.Client.CreateCourse(client.Ctx, &course.CreateCourseRequest{
		Title:       requestPayload.Title,
		Description: requestPayload.Description,
		CreatorId:   requestPayload.CreatorId,
	})

	if err != nil {
		log.Error("Failed to create course", "title", requestPayload.Title, "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully created course", "title", requestPayload.Title)

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Successfully created course titled %s", requestPayload.Title),
	}

	utils.WriteJSON(w, http.StatusCreated, responsePayload)

	log.Debug("Terminating CreateCourse request handling")
}

func (api *CourseService) HandleGetAllCourses(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling GetAllCourses request....")

	client, err := api.getCourseServiceClient()
	if err != nil {
		log.Error("Failed to get gRPC client", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	log.Info("Fetching all courses")
	res, err := client.Client.GetAllCourses(client.Ctx, &course.GetAllCoursesRequest{})
	if err != nil {
		log.Error("Failed to fetch all courses", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully fetched all courses", "count", len(res.Courses))

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched all courses",
		Data:    res.Courses,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)

	log.Debug("Terminating GetAllCourses request handling")
}

func (api *CourseService) HandleGetCourseByTitle(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling GetCourseByTitle request....")

	client, err := api.getCourseServiceClient()
	if err != nil {
		log.Error("Failed to get gRPC client", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	title := chi.URLParam(r, "title")

	log.Info("Fetching course by title", "title", title)
	c, err := client.Client.GetCourse(client.Ctx, &course.GetCourseRequest{
		Title: title,
	})

	if err != nil {
		log.Error("Failed to fetch course by title", "title", title, "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully fetched course", "title", c.Course.Title)

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched course",
		Data:    c.Course,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)

	log.Debug("Terminating GetCourseByTitle request handling")
}

func (api *CourseService) HandleUpdateCourseByTitle(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling UpdateCourseByTitle request....")

	var requestPayload Course
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	client, err := api.getCourseServiceClient()
	if err != nil {
		log.Error("Failed to get gRPC client", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	convertedModules := make([]*course.Module, len(requestPayload.Modules))
	for i, m := range requestPayload.Modules {
		convertedTasks := make([]*course.Task, len(m.Tasks))
		for j, t := range m.Tasks {
			convertedTasks[j] = &course.Task{
				Id:   t.Id,
				Task: t.Task,
				Type: t.Type,
				Xp:   t.Xp,
			}
		}
		convertedModules[i] = &course.Module{
			Id:    m.Id,
			Title: m.Title,
			Tasks: convertedTasks,
		}
	}

	log.Info("Updating course", "title", requestPayload.Title)
	_, err = client.Client.UpdateCourse(client.Ctx, &course.UpdateCourseRequest{
		Title:       requestPayload.Title,
		Description: requestPayload.Description,
		CreatorId:   requestPayload.CreatorId,
		Likes:       requestPayload.Likes,
		Students:    requestPayload.Students,
		Topics:      requestPayload.Topics,
		Modules:     convertedModules,
	})

	if err != nil {
		log.Error("Failed to update course", "title", requestPayload.Title, "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully updated course", "title", requestPayload.Title)

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Successfully updated course titled %s", requestPayload.Title),
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)

	log.Debug("Terminating UpdateCourseByTitle request handling")
}

func (api *CourseService) HandleDeleteCourseByTitle(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling DeleteCourseByTitle request....")

	title := chi.URLParam(r, "title")

	client, err := api.getCourseServiceClient()
	if err != nil {
		log.Error("Failed to get gRPC client", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	log.Info("Deleting course by title", "title", title)
	_, err = client.Client.DeleteCourse(client.Ctx, &course.DeleteCourseRequest{
		Title: title,
	})

	if err != nil {
		log.Error("Failed to delete course by title", "title", title, "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully deleted course", "title", title)

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully deleted course",
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)

	log.Debug("Terminating DeleteCourseByTitle request handling")
}
