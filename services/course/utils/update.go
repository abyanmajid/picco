package utils

import (
	"time"

	"github.com/abyanmajid/codemore.io/services/course/models"
	"github.com/abyanmajid/codemore.io/services/course/proto/course"
)

func UpdateCourseFields(course *models.Course, req *course.UpdateCourseRequest) {
	if req.GetDescription() != "" {
		course.Description = req.GetDescription()
	}
	if req.GetCreatorId() != "" {
		course.CreatorID = req.GetCreatorId()
	}
	if req.GetLikes() != 0 {
		course.Likes = req.GetLikes()
	}
	if len(req.GetStudents()) > 0 {
		course.Students = req.GetStudents()
	}
	if len(req.GetTopics()) > 0 {
		course.Topics = req.GetTopics()
	}
	if len(req.GetModules()) > 0 {
		course.Modules = ProtoDecodeModules(req.GetModules())
	}

	course.UpdatedAt = time.Now().String()
}
