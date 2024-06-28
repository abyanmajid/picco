package main

import (
	"net/http"

	"github.com/abyanmajid/codemore.io/services/broker/proto/course"
	"github.com/go-chi/chi/v5"
)

func (api *Service) HandleCreateCourse(w http.ResponseWriter, r *http.Request) {
	api.Log.Info("Handling create course request")

	var requestPayload Course
	err := api.readJSON(w, r, &requestPayload)
	if err != nil {
		api.Log.Error("Failed to read JSON from request", "error", err)
		api.errorJSON(w, err)
		return
	}

	api.Log.Info("Successfully read request payload", "payload", requestPayload)

	client, err := api.getCourseServiceClient()
	if err != nil {
		api.Log.Error("Failed to get course service client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	api.Log.Info("Creating course with course service client", "title", requestPayload.Title, "creator", requestPayload.Creator)

	c, err := client.Client.CreateCourse(client.Ctx, &course.CreateCourseRequest{
		Title:       requestPayload.Title,
		Description: requestPayload.Description,
		Creator:     requestPayload.Creator,
	})

	if err != nil {
		api.Log.Error("Failed to create course", "error", err)
		api.errorJSON(w, err)
		return
	}

	api.Log.Info("Successfully created course", "title", c.Course.Title)

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully created course",
		Data:    c.Course,
	}

	api.writeJSON(w, http.StatusCreated, responsePayload)
	api.Log.Info("Response sent", "status", http.StatusCreated)
}

func (api *Service) HandleGetAllCourses(w http.ResponseWriter, r *http.Request) {
	client, err := api.getCourseServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	res, err := client.Client.GetAllCourses(client.Ctx, &course.GetAllCoursesRequest{})
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully fetched all courses",
		Data:    res.Courses,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleGetCourseByTitle(w http.ResponseWriter, r *http.Request) {
	client, err := api.getCourseServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	c, err := client.Client.GetCourse(client.Ctx, &course.GetCourseRequest{
		Title: chi.URLParam(r, "title"),
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully fetched course",
		Data:    c.Course,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleUpdateCourseByTitle(w http.ResponseWriter, r *http.Request) {
	var requestPayload Course
	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getCourseServiceClient()
	if err != nil {
		api.errorJSON(w, err)
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
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully updated course",
		Data:    c.Course,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleDeleteCourseByTitle(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")

	client, err := api.getCourseServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	_, err = client.Client.DeleteCourse(client.Ctx, &course.DeleteCourseRequest{
		Title: title,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully deleted course",
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}
