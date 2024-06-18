package repositories

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-journal/internal/app/repositories/entities"
)

type Repository interface {
	GetStudentMarks(ctx context.Context, studyPlaceId uuid.UUID, lessonIds []uuid.UUID, studentId uuid.UUID) ([]entities.Mark, error)
	GetLessonsMarks(ctx context.Context, studyPlaceId uuid.UUID, lessonIds []uuid.UUID) ([]entities.Mark, error)
	AddMark(ctx context.Context, mark entities.Mark) error
	DeleteMark(ctx context.Context, studyPlaceId uuid.UUID, lessonId uuid.UUID, teacherId uuid.UUID, markId uuid.UUID) error
	EditMark(ctx context.Context, mark entities.Mark) error

	GetLessonInfoByLessonId(ctx context.Context, studyPlaceId uuid.UUID, lessonId uuid.UUID) (entities.LessonInfo, error)
	AddLessonInfo(ctx context.Context, lessonInfo entities.LessonInfo) error
	DeleteLessonInfo(ctx context.Context, studyPlaceId uuid.UUID, lessonId uuid.UUID, teacherId uuid.UUID, id uuid.UUID) error
	EditLessonInfo(ctx context.Context, lessonInfo entities.LessonInfo) error

	GetStudentAbsences(ctx context.Context, studyPlaceId uuid.UUID, lessonIds []uuid.UUID, studentId uuid.UUID) ([]entities.Absence, error)
	GetLessonsAbsences(ctx context.Context, studyPlaceId uuid.UUID, lessonIds []uuid.UUID) ([]entities.Absence, error)
	AddAbsence(ctx context.Context, absence entities.Absence) error
	DeleteAbsence(ctx context.Context, studyPlaceId uuid.UUID, lessonId uuid.UUID, teacherId uuid.UUID, markId uuid.UUID) error
	EditAbsence(ctx context.Context, mark entities.Absence) error
}

type repository struct {
	database *gocql.Session
}

func New(database *gocql.Session) Repository {
	return &repository{
		database: database,
	}
}
