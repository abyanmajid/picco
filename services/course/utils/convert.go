package utils

import (
	"github.com/abyanmajid/codemore.io/services/course/models"
	"github.com/abyanmajid/codemore.io/services/course/proto/course"
)

func ProtoEncodeCourse(modelCourse models.Course) *course.Course {
	return &course.Course{
		Id:          modelCourse.ID,
		Title:       modelCourse.Title,
		Description: modelCourse.Description,
		CreatorId:   modelCourse.CreatorID,
		Likes:       modelCourse.Likes,
		Students:    modelCourse.Students,
		Topics:      modelCourse.Topics,
		Modules:     ProtoEncodeModules(modelCourse.Modules),
		UpdatedAt:   modelCourse.UpdatedAt,
		CreatedAt:   modelCourse.CreatedAt,
	}
}

func ProtoEncodeModules(modelModules []models.Module) []*course.Module {
	var protoModules []*course.Module
	for _, modelModule := range modelModules {
		protoModules = append(protoModules, &course.Module{
			Id:    modelModule.ID,
			Title: modelModule.Title,
			Tasks: ProtoEncodeTasks(modelModule.Tasks),
		})
	}
	return protoModules
}

func ProtoEncodeTasks(modelTasks []models.Task) []*course.Task {
	var protoTasks []*course.Task
	for _, modelTask := range modelTasks {
		protoTasks = append(protoTasks, &course.Task{
			Id:   modelTask.ID,
			Task: modelTask.Task,
			Type: modelTask.Type,
			Xp:   modelTask.XP,
		})
	}
	return protoTasks
}

func ProtoDecodeModules(protoModules []*course.Module) []models.Module {
	var modelModules []models.Module
	for _, protoModule := range protoModules {
		modelModule := models.Module{
			ID:    protoModule.GetId(),
			Title: protoModule.GetTitle(),
			Tasks: ProtoDecodeTasks(protoModule.GetTasks()),
		}
		modelModules = append(modelModules, modelModule)
	}
	return modelModules
}

func ProtoDecodeTasks(protoTasks []*course.Task) []models.Task {
	var modelTasks []models.Task
	for _, protoTask := range protoTasks {
		modelTask := models.Task{
			ID:   protoTask.GetId(),
			Task: protoTask.GetTask(),
			Type: protoTask.GetType(),
			XP:   protoTask.GetXp(),
		}
		modelTasks = append(modelTasks, modelTask)
	}
	return modelTasks
}
