package main

import (
	"context"
	"fmt"
	"time"

	course "github.com/abyanmajid/codemore.io/services/course/proto/course"
	"go.mongodb.org/mongo-driver/bson"
)

func convertModulesToPb(modules []Module) []*course.Module {
	var pbModules []*course.Module
	for _, m := range modules {
		pbModules = append(pbModules, &course.Module{
			Id:    m.Id,
			Title: m.Title,
			Tasks: convertTasksToPb(m.Tasks),
		})
	}
	return pbModules
}

func convertTasksToPb(tasks []Task) []*course.Task {
	var pbTasks []*course.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &course.Task{
			Id:   t.Id,
			Task: t.Task,
			Type: t.Type,
			Xp:   t.Xp,
		})
	}
	return pbTasks
}

func (api *Service) CreateCourse(ctx context.Context, req *course.CreateCourseRequest) (*course.CreateCourseResponse, error) {
	api.Log.Info("Creating course", "title", req.GetTitle(), "creator", req.GetCreator())

	c := Course{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Creator:     req.GetCreator(),
		Likes:       0,
		Students:    []string{},
		Topics:      []string{},
		Modules:     []Module{},
		UpdatedAt:   time.Now().String(),
		CreatedAt:   time.Now().String(),
	}

	collection := api.Mongo.Database("db").Collection("courses")
	_, err := collection.InsertOne(ctx, c)
	if err != nil {
		api.Log.Error("Failed to insert course into MongoDB", "error", err)
		return nil, err
	}

	api.Log.Info("Successfully inserted course into MongoDB", "title", c.Title, "creator", c.Creator)

	createdCourse := &course.Course{
		Title:       c.Title,
		Description: c.Description,
		Creator:     c.Creator,
		Likes:       c.Likes,
		Students:    c.Students,
		Topics:      c.Topics,
		Modules:     []*course.Module{},
		UpdatedAt:   c.UpdatedAt,
		CreatedAt:   c.CreatedAt,
	}

	api.Log.Info("Successfully created course response", "title", createdCourse.Title, "creator", createdCourse.Creator)

	return &course.CreateCourseResponse{
		Course: createdCourse,
	}, nil
}

func (api *Service) GetAllCourses(ctx context.Context, req *course.GetAllCoursesRequest) (*course.GetAllCoursesResponse, error) {
	collection := api.Mongo.Database("db").Collection("courses")

	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var courses []Course
	for cursor.Next(ctx) {
		var course Course
		if err := cursor.Decode(&course); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	var pbCourses []*course.Course
	for _, c := range courses {
		pbCourse := &course.Course{
			Title:       c.Title,
			Description: c.Description,
			Creator:     c.Creator,
			Likes:       c.Likes,
			Students:    c.Students,
			Topics:      c.Topics,
			Modules:     convertModulesToPb(c.Modules),
			UpdatedAt:   c.UpdatedAt,
			CreatedAt:   c.CreatedAt,
		}
		pbCourses = append(pbCourses, pbCourse)
	}

	res := &course.GetAllCoursesResponse{
		Courses: pbCourses,
	}

	return res, nil
}

func (api *Service) GetCourse(ctx context.Context, req *course.GetCourseRequest) (*course.GetCourseResponse, error) {
	collection := api.Mongo.Database("db").Collection("courses")

	filter := bson.M{
		"title": req.GetTitle(),
	}
	doc := collection.FindOne(ctx, filter)

	var c Course
	if err := doc.Decode(&c); err != nil {
		return nil, err
	}

	pbCourse := &course.Course{
		Title:       c.Title,
		Description: c.Description,
		Creator:     c.Creator,
		Likes:       c.Likes,
		Students:    c.Students,
		Topics:      c.Topics,
		Modules:     convertModulesToPb(c.Modules),
		UpdatedAt:   c.UpdatedAt,
		CreatedAt:   c.CreatedAt,
	}

	return &course.GetCourseResponse{
		Course: pbCourse,
	}, nil
}

func (api *Service) UpdateCourse(ctx context.Context, req *course.UpdateCourseRequest) (*course.UpdateCourseResponse, error) {
	collection := api.Mongo.Database("db").Collection("courses")

	updateFields := bson.M{}

	updateFields["title"] = req.GetTitle()

	if req.GetDescription() != "" {
		updateFields["description"] = req.GetDescription()
	}
	if req.GetCreator() != "" {
		updateFields["creator"] = req.GetCreator()
	}
	if req.GetLikes() != 0 {
		updateFields["likes"] = req.GetLikes()
	}
	if len(req.GetStudents()) > 0 {
		updateFields["students"] = req.GetStudents()
	}
	if len(req.GetTopics()) > 0 {
		updateFields["topics"] = req.GetTopics()
	}
	if len(req.GetModules()) > 0 {
		updateFields["modules"] = req.GetModules()
	}
	updateFields["updated_at"] = time.Now().String()

	update := bson.M{"$set": updateFields}
	filter := bson.M{"title": req.GetTitle()}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no matching course found")
	}

	doc := collection.FindOne(ctx, filter)
	var updatedCourse Course
	if err := doc.Decode(&updatedCourse); err != nil {
		return nil, err
	}

	convertedModules := make([]*course.Module, len(updatedCourse.Modules))
	for i, m := range updatedCourse.Modules {
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

	return &course.UpdateCourseResponse{
		Course: &course.Course{
			Title:       updatedCourse.Title,
			Description: updatedCourse.Description,
			Creator:     updatedCourse.Creator,
			Likes:       updatedCourse.Likes,
			Topics:      updatedCourse.Topics,
			Modules:     convertedModules,
			UpdatedAt:   updatedCourse.UpdatedAt,
			CreatedAt:   updatedCourse.CreatedAt,
		},
	}, nil
}

func (api *Service) DeleteCourse(ctx context.Context, req *course.DeleteCourseRequest) (*course.DeleteCourseResponse, error) {
	collection := api.Mongo.Database("db").Collection("courses")

	filter := bson.M{"title": req.GetTitle()}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, err
	}

	return &course.DeleteCourseResponse{}, nil
}
