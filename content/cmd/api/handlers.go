package main

import (
	"context"
	"fmt"
	"time"

	"github.com/abyanmajid/codemore.io/content/proto/content"
	"github.com/abyanmajid/codemore.io/content/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (api *Service) CreateCourse(ctx context.Context, req *content.CreateCourseRequest) (*content.CreateCourseResponse, error) {
	course := Course{
		Title:     req.GetTitle(),
		Creator:   req.GetCreator(),
		Likes:     0,
		Topics:    []string{},
		Modules:   []Module{},
		UpdatedAt: time.Now().String(),
		CreatedAt: time.Now().String(),
	}

	collection := api.Mongo.Database("db").Collection("courses")
	doc, err := collection.InsertOne(ctx, course)
	if err != nil {
		return nil, err
	}

	insertedId, err := utils.ConvertToObjectIDString(doc.InsertedID)
	if err != nil {
		return nil, err
	}

	return &content.CreateCourseResponse{
		Course: &content.Course{
			Id:        insertedId,
			Title:     course.Title,
			Creator:   course.Creator,
			Likes:     course.Likes,
			Topics:    course.Topics,
			Modules:   []*content.Module{},
			UpdatedAt: course.UpdatedAt,
			CreatedAt: course.CreatedAt,
		},
	}, nil
}

func (api *Service) GetAllCourses(ctx context.Context, req *content.GetAllCoursesRequest) (*content.GetAllCoursesResponse, error) {
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

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	res := &content.GetAllCoursesResponse{}
	for _, c := range courses {
		convertedModules := make([]*content.Module, len(c.Modules))
		for i, m := range c.Modules {
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

		testCase := &content.Course{
			Id:        c.Id,
			Title:     c.Title,
			Creator:   c.Creator,
			Likes:     c.Likes,
			Topics:    c.Topics,
			Modules:   convertedModules,
			UpdatedAt: c.UpdatedAt,
			CreatedAt: c.CreatedAt,
		}
		res.Courses = append(res.Courses, testCase)
	}

	return res, nil
}

func (api *Service) GetCourse(ctx context.Context, req *content.GetCourseRequest) (*content.GetCourseResponse, error) {
	collection := api.Mongo.Database("db").Collection("courses")

	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	doc := collection.FindOne(ctx, filter)

	var course Course
	if err := doc.Decode(&course); err != nil {
		return nil, err
	}

	// Convert course.Modules to []*content.Module
	convertedModules := make([]*content.Module, len(course.Modules))
	for i, m := range course.Modules {
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

	return &content.GetCourseResponse{
		Course: &content.Course{
			Id:        course.Id,
			Title:     course.Title,
			Creator:   course.Creator,
			Likes:     course.Likes,
			Topics:    course.Topics,
			Modules:   convertedModules,
			UpdatedAt: course.UpdatedAt,
			CreatedAt: course.CreatedAt,
		},
	}, nil
}

func (api *Service) UpdateCourse(ctx context.Context, req *content.UpdateCourseRequest) (*content.UpdateCourseResponse, error) {
	collection := api.Mongo.Database("db").Collection("courses")

	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	updateFields := bson.M{}
	if req.GetTitle() != "" {
		updateFields["title"] = req.GetTitle()
	}
	if req.GetCreator() != "" {
		updateFields["creator"] = req.GetCreator()
	}
	if req.GetLikes() != 0 {
		updateFields["likes"] = req.GetLikes()
	}
	if len(req.GetTopics()) > 0 {
		updateFields["topics"] = req.GetTopics()
	}
	if len(req.GetModules()) > 0 {
		// Convert req.GetModules() to a format that MongoDB understands
		modules := make([]Module, len(req.GetModules()))
		for i, m := range req.GetModules() {
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
		updateFields["modules"] = modules
	}
	updateFields["updated_at"] = time.Now().String()

	update := bson.M{"$set": updateFields}
	filter := bson.M{"_id": objectID}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no matching course found")
	}

	// Retrieve the updated document to include in the response
	doc := collection.FindOne(ctx, filter)
	var updatedCourse Course
	if err := doc.Decode(&updatedCourse); err != nil {
		return nil, err
	}

	// Convert updatedCourse.Modules to []*content.Module for response
	convertedModules := make([]*content.Module, len(updatedCourse.Modules))
	for i, m := range updatedCourse.Modules {
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

	return &content.UpdateCourseResponse{
		Course: &content.Course{
			Id:        updatedCourse.Id,
			Title:     updatedCourse.Title,
			Creator:   updatedCourse.Creator,
			Likes:     updatedCourse.Likes,
			Topics:    updatedCourse.Topics,
			Modules:   convertedModules,
			UpdatedAt: updatedCourse.UpdatedAt,
			CreatedAt: updatedCourse.CreatedAt,
		},
	}, nil
}

func (api *Service) DeleteCourse(ctx context.Context, req *content.DeleteCourseRequest) (*content.DeleteCourseResponse, error) {
	collection := api.Mongo.Database("db").Collection("courses")
	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &content.DeleteCourseResponse{}, nil
}
