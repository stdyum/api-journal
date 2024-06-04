package repositories

import (
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-journal/internal/app/repositories/entities"
)

func (r *repository) scanMark(row databases.Scan) (mark entities.Mark, err error) {
	var id, studyPlaceId, studentId, teacherId, lessonId gocql.UUID

	err = row.Scan(
		&id,
		&studyPlaceId,
		&mark.Mark,
		&studentId,
		&teacherId,
		&lessonId,
	)

	mark.ID = uuid.UUID(id)
	mark.StudyPlaceId = uuid.UUID(studyPlaceId)
	mark.StudentId = uuid.UUID(studentId)
	mark.TeacherId = uuid.UUID(teacherId)
	mark.LessonId = uuid.UUID(lessonId)

	return
}

func (r *repository) scanLessonInfo(row databases.Scan) (lessonInfo entities.LessonInfo, err error) {
	var id, studyPlaceId, lessonId, teacherId gocql.UUID

	err = row.Scan(
		&id,
		&studyPlaceId,
		&lessonId,
		&teacherId,
		&lessonInfo.Title,
		&lessonInfo.Description,
		&lessonInfo.Homework,
		&lessonInfo.Type,
	)

	lessonInfo.ID = uuid.UUID(id)
	lessonInfo.StudyPlaceId = uuid.UUID(studyPlaceId)
	lessonInfo.LessonId = uuid.UUID(lessonId)
	lessonInfo.TeacherId = uuid.UUID(teacherId)

	return
}
