package handlers

import (
	"net/http"

	log "log/slog"

	"github.com/abyanmajid/codemore.io/services/course/models"
	"github.com/abyanmajid/codemore.io/services/course/repositories"
	"github.com/abyanmajid/codemore.io/services/course/utils"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourseService struct {
	Collection *mongo.Collection
}

func (s *CourseService) HandleCreateCourse(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling CreateCourse request....")

	var requestPayload models.WriteCourse
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Creating new course", "title", requestPayload.Title)
	repo := repositories.CourseRepository{Collection: s.Collection}
	res, err := repo.CreateCourse(&requestPayload)
	if err != nil {
		log.Error("Failed to create course", "title", requestPayload.Title, "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully created course", "course_id", res.InsertedID)
	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully created course",
		Data:    res.InsertedID,
	}

	utils.WriteJSON(w, http.StatusCreated, responsePayload)
	log.Debug("Terminating CreateCourse request handling")
}

func (s *CourseService) HandleGetAllCourses(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling GetAllCourses request....")

	repo := repositories.CourseRepository{Collection: s.Collection}
	courses, err := repo.GetAllCourses()
	if err != nil {
		log.Error("Failed to fetch all courses", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully fetched all courses", "count", len(courses))
	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched all courses",
		Data:    courses,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	log.Debug("Terminating GetAllCourses request handling")
}

func (s *CourseService) HandleGetCourseByTitle(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling GetCourseByTitle request....")

	title := chi.URLParam(r, "title")
	log.Debug("Fetched URL parameter", "title", title)

	repo := repositories.CourseRepository{Collection: s.Collection}
	course, err := repo.GetCourseByTitle(title)
	if err != nil {
		log.Error("Failed to fetch course by title", "title", title, "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully fetched course", "course_id", course.ID)
	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched course",
		Data:    course,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	log.Debug("Terminating GetCourseByTitle request handling")
}

func (s *CourseService) HandleUpdateCourseByTitle(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling UpdateCourseByTitle request....")

	title := chi.URLParam(r, "title")
	log.Debug("Fetched URL parameter", "title", title)

	var requestPayload models.WriteCourse
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Updating course", "title", title)
	repo := repositories.CourseRepository{Collection: s.Collection}
	result, err := repo.UpdateCourseByTitle(title, &requestPayload)
	if err != nil {
		log.Error("Failed to update course", "title", title, "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully updated course", "modified_count", result.ModifiedCount)
	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully updated course",
		Data:    result.ModifiedCount,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	log.Debug("Terminating UpdateCourseByTitle request handling")
}

func (s *CourseService) HandleDeleteCourseByTitle(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling DeleteCourseByTitle request....")

	title := chi.URLParam(r, "title")
	log.Debug("Fetched URL parameter", "title", title)

	repo := repositories.CourseRepository{Collection: s.Collection}
	result, err := repo.DeleteCourseByTitle(title)
	if err != nil {
		log.Error("Failed to delete course by title", "title", title, "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	log.Info("Successfully deleted course", "deleted_count", result.DeletedCount)
	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully deleted course",
		Data:    result.DeletedCount,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	log.Debug("Terminating DeleteCourseByTitle request handling")
}
