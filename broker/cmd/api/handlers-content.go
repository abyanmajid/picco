package main

import (
	"net/http"
	"time"

	"github.com/abyanmajid/codemore.io/broker/proto/content"
	"github.com/go-chi/chi/v5"
)

func (api *Service) HandleCreateCourse(w http.ResponseWriter, r *http.Request) {
	var requestPayload Course
	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getContentServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	c, err := client.Client.CreateCourse(client.Ctx, &content.CreateCourseRequest{
		Title:   requestPayload.Title,
		Creator: requestPayload.Creator,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully created course:",
		Data:    c,
	}

	api.writeJSON(w, http.StatusCreated, responsePayload)
}

func (api *Service) HandleGetAllCourses(w http.ResponseWriter, r *http.Request) {
	client, err := api.getContentServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	res, err := client.Client.GetAllCourses(client.Ctx, &content.GetAllCoursesRequest{})
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	// Convert res.Courses to JSON-friendly format
	courses := make([]Course, len(res.Courses))
	for i, c := range res.Courses {
		modules := make([]Module, len(c.Modules))
		for j, m := range c.Modules {
			tasks := make([]Task, len(m.Tasks))
			for k, t := range m.Tasks {
				tasks[k] = Task{
					Id:    t.Id,
					Title: t.Title,
					Mdx:   t.Mdx,
				}
			}
			modules[j] = Module{
				Id:    m.Id,
				Title: m.Title,
				Tasks: tasks,
			}
		}
		courses[i] = Course{
			Id:        c.Id,
			Title:     c.Title,
			Creator:   c.Creator,
			Likes:     c.Likes,
			Topics:    c.Topics,
			Modules:   modules,
			UpdatedAt: c.UpdatedAt,
			CreatedAt: c.CreatedAt,
		}
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully fetched all courses",
		Data:    courses,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleGetCourseById(w http.ResponseWriter, r *http.Request) {
	client, err := api.getContentServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	c, err := client.Client.GetCourse(client.Ctx, &content.GetCourseRequest{
		Id: chi.URLParam(r, "id"),
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	// Convert response Course to JSON-friendly format
	modules := make([]Module, len(c.Course.Modules))
	for i, m := range c.Course.Modules {
		tasks := make([]Task, len(m.Tasks))
		for j, t := range m.Tasks {
			tasks[j] = Task{
				Id:    t.Id,
				Title: t.Title,
				Mdx:   t.Mdx,
			}
		}
		modules[i] = Module{
			Id:    m.Id,
			Title: m.Title,
			Tasks: tasks,
		}
	}
	course := Course{
		Id:        c.Course.Id,
		Title:     c.Course.Title,
		Creator:   c.Course.Creator,
		Likes:     c.Course.Likes,
		Topics:    c.Course.Topics,
		Modules:   modules,
		UpdatedAt: c.Course.UpdatedAt,
		CreatedAt: c.Course.CreatedAt,
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully fetched course",
		Data:    course,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleUpdateCourseById(w http.ResponseWriter, r *http.Request) {
	var requestPayload Course
	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getContentServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	// Convert requestPayload.Modules to []*content.Module
	convertedModules := make([]*content.Module, len(requestPayload.Modules))
	for i, m := range requestPayload.Modules {
		convertedTasks := make([]*content.Task, len(m.Tasks))
		for j, t := range m.Tasks {
			convertedTasks[j] = &content.Task{
				Id:    t.Id,
				Title: t.Title,
				Mdx:   t.Mdx,
			}
		}
		convertedModules[i] = &content.Module{
			Id:    m.Id,
			Title: m.Title,
			Tasks: convertedTasks,
		}
	}

	c, err := client.Client.UpdateCourse(client.Ctx, &content.UpdateCourseRequest{
		Id:        chi.URLParam(r, "id"),
		Title:     requestPayload.Title,
		Creator:   requestPayload.Creator,
		Likes:     requestPayload.Likes,
		Topics:    requestPayload.Topics,
		Modules:   convertedModules,
		UpdatedAt: time.Now().String(),
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

func (api *Service) HandleDeleteCourseById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	client, err := api.getContentServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	_, err = client.Client.DeleteCourse(client.Ctx, &content.DeleteCourseRequest{
		Id: id,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully deleted course #" + id,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}
