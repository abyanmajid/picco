package handlers

import (
	"context"
	"fmt"
	"net"
	"time"

	l "log"
	log "log/slog"

	"github.com/abyanmajid/codemore.io/services/course/models"
	"github.com/abyanmajid/codemore.io/services/course/proto/course"
	"github.com/abyanmajid/codemore.io/services/course/repositories"
	"github.com/abyanmajid/codemore.io/services/course/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type CourseServiceParameters struct {
	Port       string
	App        string
	Client     *mongo.Client
	Database   string
	Collection string
}

func ListenAndServeCourse(params CourseServiceParameters) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", params.Port))
	if err != nil {
		l.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	course.RegisterCourseServiceServer(s, &Service{
		Collection: params.Client.Database(params.Database).Collection(params.Collection),
	})

	l.Printf("gRPC Server started on port %s", params.Port)

	if err := s.Serve(lis); err != nil {
		l.Fatalf("Failed to listen for gRPC: %v", err)
	}
}

func (s *Service) CreateCourse(ctx context.Context, req *course.CreateCourseRequest) (*course.CreateCourseResponse, error) {
	log.Debug("Performing CreateCourse....")

	c := models.Course{
		ID:          uuid.New().String(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		CreatorID:   req.GetCreatorId(),
		Likes:       0,
		Students:    []string{},
		Topics:      []string{},
		Modules:     []models.Module{},
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	log.Debug("Created course model for new course", "course_id", c.ID)

	repo := repositories.CourseRepository{Collection: s.Collection}

	log.Debug("Attained CourseRepository")

	log.Info("Creating a new course", "course_id", c.ID)
	_, err := repo.CreateCourse(&c)
	if err != nil {
		log.Error("Failed to create course", "course_id", c.ID, "error", err)
		return nil, err
	}
	log.Info("Successfully created course", "course_id", c.ID)

	log.Debug("Terminating CreateCourse...")

	return &course.CreateCourseResponse{
		CreatedCount: 1,
	}, nil
}

func (s *Service) GetAllCourses(ctx context.Context, req *course.GetAllCoursesRequest) (*course.GetAllCoursesResponse, error) {
	log.Debug("Performing GetAllCourses....")

	repo := repositories.CourseRepository{Collection: s.Collection}

	log.Debug("Attained CourseRepository")

	results, err := repo.GetAllCourses()
	if err != nil {
		log.Error("Failed to get all courses", "error", err)
		return nil, err
	}

	log.Debug("Fetched courses from repository", "count", len(results))

	var protoCourses []*course.Course
	for _, modelCourse := range results {
		protoCourses = append(protoCourses, utils.ProtoEncodeCourse(modelCourse))
	}

	log.Debug("Converted courses to proto format", "count", len(protoCourses))

	log.Debug("Terminating GetAllCourses...")

	return &course.GetAllCoursesResponse{
		Courses: protoCourses,
	}, nil
}

type Service struct {
	course.UnimplementedCourseServiceServer
	Collection *mongo.Collection
}

func (s *Service) GetCourse(ctx context.Context, req *course.GetCourseRequest) (*course.GetCourseResponse, error) {
	log.Debug("Performing GetCourseByTitle....")

	repo := repositories.CourseRepository{Collection: s.Collection}

	log.Debug("Attained CourseRepository")

	c, err := repo.GetCourseByTitle(req.GetTitle())
	if err != nil {
		log.Error("Failed to get course by title", "title", req.GetTitle(), "error", err)
		return nil, err
	}

	log.Debug("Fetched course from repository", "course_id", c.ID)

	protoCourse := utils.ProtoEncodeCourse(*c)

	log.Debug("Converted course to proto format", "course_id", protoCourse.Id)

	log.Debug("Terminating GetCourseByTitle...")

	return &course.GetCourseResponse{Course: protoCourse}, nil
}

func (s *Service) UpdateCourse(ctx context.Context, req *course.UpdateCourseRequest) (*course.UpdateCourseResponse, error) {
	log.Debug("Performing UpdateCourseByTitle....")

	repo := repositories.CourseRepository{Collection: s.Collection}

	log.Debug("Attained CourseRepository")

	// Fetch the existing course
	existingCourse, err := repo.GetCourseByTitle(req.GetTitle())
	if err != nil {
		log.Error("Failed to get course by title", "title", req.GetTitle(), "error", err)
		return nil, err
	}

	log.Debug("Fetched existing course", "course_id", existingCourse.ID)

	// Update fields conditionally
	utils.UpdateCourseFields(existingCourse, req)

	log.Debug("Updated course fields", "course_id", existingCourse.ID)

	res, err := repo.UpdateCourseByTitle(req.GetTitle(), existingCourse)
	if err != nil {
		log.Error("Failed to update course by title", "course_id", existingCourse.ID, "error", err)
		return nil, err
	}

	log.Info("Successfully updated course", "course_id", existingCourse.ID)

	log.Debug("Terminating UpdateCourseByTitle...")

	return &course.UpdateCourseResponse{
		UpdatedCount: res.ModifiedCount,
	}, nil
}

func (s *Service) DeleteCourse(ctx context.Context, req *course.DeleteCourseRequest) (*course.DeleteCourseResponse, error) {
	log.Debug("Performing DeleteCourseByTitle....")

	repo := repositories.CourseRepository{Collection: s.Collection}

	log.Debug("Attained CourseRepository")

	log.Info("Deleting course by title", "title", req.GetTitle())

	result, err := repo.DeleteCourseByTitle(req.GetTitle())
	if err != nil {
		log.Error("Failed to delete course by title", "title", req.GetTitle(), "error", err)
		return nil, err
	}

	log.Info("Successfully deleted course", "title", req.GetTitle())

	log.Debug("Terminating DeleteCourseByTitle...")

	return &course.DeleteCourseResponse{
		DeletedCount: result.DeletedCount,
	}, nil
}
