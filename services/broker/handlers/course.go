package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Creator     string   `json:"creator"`
	Likes       int32    `json:"likes"`
	Students    []string `json:"students"`
	Topics      []string `json:"topics"`
	Modules     []Module `json:"modules"`
	UpdatedAt   string   `json:"updated_at"`
	CreatedAt   string   `json:"created_at"`
}

type Module struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id   int32  `json:"id"`
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

	var requestPayload Course
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	client, err := api.getCourseServiceClient()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	c, err := client.Client.CreateCourse(client.Ctx, &course.CreateCourseRequest{
		Title:       requestPayload.Title,
		Description: requestPayload.Description,
		Creator:     requestPayload.Creator,
	})

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully created course",
		Data:    c.Course,
	}

	utils.WriteJSON(w, http.StatusCreated, responsePayload)
}

func (api *CourseService) HandleGetAllCourses(w http.ResponseWriter, r *http.Request) {
	client, err := api.getCourseServiceClient()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	res, err := client.Client.GetAllCourses(client.Ctx, &course.GetAllCoursesRequest{})
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched all courses",
		Data:    res.Courses,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}

func (api *CourseService) HandleGetCourseByTitle(w http.ResponseWriter, r *http.Request) {
	client, err := api.getCourseServiceClient()
	if err != nil {
		utils.ErrorJSON(w, err)
		fmt.Println("Failed to get gRPC client:", err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	title := chi.URLParam(r, "title")

	// Make gRPC call with timeout
	c, err := client.Client.GetCourse(client.Ctx, &course.GetCourseRequest{
		Title: title,
	})

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched course",
		Data:    c.Course,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}

func (api *CourseService) HandleUpdateCourseByTitle(w http.ResponseWriter, r *http.Request) {
	var requestPayload Course
	err := utils.ReadJSON(w, r, &requestPayload)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	client, err := api.getCourseServiceClient()
	if err != nil {
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

	c, err := client.Client.UpdateCourse(client.Ctx, &course.UpdateCourseRequest{
		Title:       requestPayload.Title,
		Description: requestPayload.Description,
		Creator:     requestPayload.Creator,
		Likes:       requestPayload.Likes,
		Students:    requestPayload.Students,
		Topics:      requestPayload.Topics,
		Modules:     convertedModules,
	})

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully updated course",
		Data:    c.Course,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}

func (api *CourseService) HandleDeleteCourseByTitle(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")

	client, err := api.getCourseServiceClient()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	_, err = client.Client.DeleteCourse(client.Ctx, &course.DeleteCourseRequest{
		Title: title,
	})

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully deleted course",
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}
