package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "log/slog"

	"github.com/abyanmajid/codemore.io/services/broker/proto/course"
	"github.com/abyanmajid/codemore.io/services/broker/utils"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
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

func (api *CourseService) HandleCreateCourse(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling CreateCourse request....")

	var requestPayload Course
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	jsonData, _ := json.Marshal(requestPayload)
	request, err := http.NewRequest("POST", api.Endpoint+"/course", bytes.NewBuffer(jsonData))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		utils.ErrorJSON(w, errors.New("response status code is not: StatusCreated"))
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Successfully created course #%s", requestPayload.Id),
	}

	log.Debug("Terminating CreateCourse request handling")

	utils.WriteJSON(w, http.StatusCreated, responsePayload)
}

func (api *CourseService) HandleGetAllCourses(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling GetAllCourses request....")

	request, err := http.NewRequest("GET", api.Endpoint+"/course", nil)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		utils.ErrorJSON(w, errors.New("response status is not: StatusOK"))
		return
	}

	var jsonFromService utils.JsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched all courses",
		Data:    jsonFromService.Data,
	}

	log.Debug("Terminating GetAllCourses request handling")

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}

func (api *CourseService) HandleGetCourseByTitle(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling GetCourseByTitle request....")

	title := chi.URLParam(r, "title")
	request, err := http.NewRequest("GET", api.Endpoint+"/course/"+title, nil)
	if err != nil {
		log.Error("Failed to create new request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		utils.ErrorJSON(w, errors.New("response status is not: StatusOK"))
		return
	}

	var jsonFromService utils.JsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched all courses",
		Data:    jsonFromService.Data,
	}

	log.Debug("Terminating GetCourseByTitle request handling")

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}

func (api *CourseService) HandleUpdateCourseByTitle(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling UpdateCourseByTitle request....")

	title := chi.URLParam(r, "title")

	var requestPayload Course
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	jsonData, _ := json.Marshal(requestPayload)
	request, err := http.NewRequest("PUT", api.Endpoint+"/course/"+title, bytes.NewReader(jsonData))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		utils.ErrorJSON(w, errors.New("response status code is not: StatusOK"))
		return
	}

	log.Info("Successfully updated course", "title", requestPayload.Title)

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Successfully updated course titled %s", requestPayload.Title),
	}

	log.Debug("Terminating UpdateCourseByTitle request handling")

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}

func (api *CourseService) HandleDeleteCourseByTitle(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling DeleteCourseByTitle request....")

	title := chi.URLParam(r, "title")
	request, err := http.NewRequest("DELETE", api.Endpoint+"/course/"+title, nil)
	if err != nil {
		log.Error("Failed to create new request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		utils.ErrorJSON(w, errors.New("response status is not: StatusOK"))
		return
	}

	log.Debug("Terminating DeleteCourseByTitle request handling")

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Successfully deleted course titled %s", title),
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}
