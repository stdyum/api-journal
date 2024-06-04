package repositories

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-journal/internal/app/repositories/entities"
)

func (r *repository) GetLessonInfoByLessonId(ctx context.Context, studyPlaceId uuid.UUID, lessonId uuid.UUID) (entities.LessonInfo, error) {
	scanner := r.database.Query(
		"SELECT id, study_place_id, lesson_id, teacher_id, title, description, homework, type FROM journal.lessons_info WHERE study_place_id = ? AND lesson_id = ? ALLOW FILTERING",
		gocql.UUID(studyPlaceId),
		gocql.UUID(lessonId),
	).
		WithContext(ctx)

	return r.scanLessonInfo(scanner)
}

func (r *repository) AddLessonInfo(ctx context.Context, lessonInfo entities.LessonInfo) error {
	return r.database.Query(`
INSERT INTO journal.lessons_info 
    (id, study_place_id, lesson_id, teacher_id, title, description, homework, type, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, dateOf(now()), dateOf(now()));`,
		gocql.UUID(lessonInfo.ID),
		gocql.UUID(lessonInfo.StudyPlaceId),
		gocql.UUID(lessonInfo.LessonId),
		gocql.UUID(lessonInfo.TeacherId),
		lessonInfo.Title,
		lessonInfo.Description,
		lessonInfo.Homework,
		lessonInfo.Type,
	).
		WithContext(ctx).
		Exec()
}

func (r *repository) DeleteLessonInfo(ctx context.Context, studyPlaceId uuid.UUID, lessonId uuid.UUID, teacherId uuid.UUID, id uuid.UUID) error {
	return r.database.Query(`
DELETE FROM journal.lessons_info WHERE study_place_id = ? AND lesson_id = ? AND id = ? AND teacher_id = ? 
`,
		gocql.UUID(studyPlaceId),
		gocql.UUID(lessonId),
		gocql.UUID(id),
		gocql.UUID(teacherId),
	).WithContext(ctx).Exec()
}

func (r *repository) EditLessonInfo(ctx context.Context, lessonInfo entities.LessonInfo) error {
	return r.database.Query(`
UPDATE journal.marks SET  
	title = ?,
	description = ?,
	homework = ?,
	type = ?,
	updated_at = dateOf(now())
WHERE 
    study_place_id = ? AND lesson_id = ? AND id = ? AND teacher_id = ? 
IF EXISTS
`,
		lessonInfo.Title,
		lessonInfo.Description,
		lessonInfo.Homework,
		lessonInfo.Type,
		gocql.UUID(lessonInfo.StudyPlaceId),
		gocql.UUID(lessonInfo.LessonId),
		gocql.UUID(lessonInfo.ID),
		gocql.UUID(lessonInfo.TeacherId),
	).WithContext(ctx).Exec()
}
