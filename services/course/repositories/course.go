package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/abyanmajid/codemore.io/services/course/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourseRepository struct {
	Collection *mongo.Collection
}

func (r *CourseRepository) CreateCourse(course *models.WriteCourse) (*mongo.InsertOneResult, error) {
	existingCourse, err := r.GetCourseByTitle(course.Title)
	if err != nil {
		return nil, err
	}

	if existingCourse != nil {
		return nil, errors.New("course with the same title already exists")
	}

	result, err := r.Collection.InsertOne(context.TODO(), course)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *CourseRepository) GetAllCourses() ([]models.ReadCourse, error) {
	results, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var courses []models.ReadCourse
	err = results.All(context.TODO(), &courses)
	if err != nil {
		return nil, fmt.Errorf("failed to decode results: %s", err.Error())
	}

	return courses, nil
}

func (r *CourseRepository) GetCourseByTitle(title string) (*models.ReadCourse, error) {
	var course models.ReadCourse

	err := r.Collection.FindOne(
		context.TODO(),
		bson.D{{Key: "title", Value: title}},
	).Decode(&course)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &course, nil
}

func (r *CourseRepository) UpdateCourseByTitle(title string, updatedCourseDetails *models.WriteCourse) (*mongo.UpdateResult, error) {
	result, err := r.Collection.UpdateOne(
		context.TODO(),
		bson.D{{Key: "title", Value: title}},
		bson.D{{Key: "$set", Value: updatedCourseDetails}},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *CourseRepository) DeleteCourseByTitle(title string) (*mongo.DeleteResult, error) {
	result, err := r.Collection.DeleteOne(
		context.TODO(),
		bson.D{{Key: "title", Value: title}},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
