package repositories

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-journal/internal/app/repositories/entities"
)

func (r *repository) GetStudentMarks(ctx context.Context, studyPlaceId uuid.UUID, lessonIds []uuid.UUID, studentId uuid.UUID) ([]entities.Mark, error) {
	scanner := r.database.Query(
		"SELECT id, study_place_id, mark, student_id, teacher_id, lesson_id FROM marks WHERE study_place_id = ? AND student_id = ? ALLOW FILTERING",
		gocql.UUID(studyPlaceId),
		//uslices.MapFunc(lessonIds, func(item uuid.UUID) gocql.UUID { return gocql.UUID(item) }),
		gocql.UUID(studentId),
	).
		WithContext(ctx).
		Iter().
		Scanner()

	return databases.ScanArray(scanner, r.scanMark)
}

func (r *repository) GetLessonsMarks(ctx context.Context, studyPlaceId uuid.UUID, lessonIds []uuid.UUID) ([]entities.Mark, error) {
	scanner := r.database.Query(
		"SELECT id, study_place_id, mark, student_id, teacher_id, lesson_id FROM marks WHERE study_place_id = ? AND lesson_id IN ? ALLOW FILTERING",
		gocql.UUID(studyPlaceId),
		uslices.MapFunc(lessonIds, func(item uuid.UUID) gocql.UUID { return gocql.UUID(item) }),
	).
		WithContext(ctx).
		Iter().
		Scanner()

	return databases.ScanArray(scanner, r.scanMark)
}

func (r *repository) AddMark(ctx context.Context, mark entities.Mark) error {
	return r.database.Query(`
INSERT INTO journal.marks 
    (id, study_place_id, mark, student_id, teacher_id, lesson_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, dateOf(now()), dateOf(now()));`,
		gocql.UUID(mark.ID),
		gocql.UUID(mark.StudyPlaceId),
		mark.Mark,
		gocql.UUID(mark.StudentId),
		gocql.UUID(mark.TeacherId),
		gocql.UUID(mark.LessonId),
	).
		WithContext(ctx).
		Exec()
}

func (r *repository) DeleteMark(ctx context.Context, studyPlaceId uuid.UUID, lessonId uuid.UUID, teacherId uuid.UUID, markId uuid.UUID) error {
	return r.database.Query(`
DELETE FROM journal.marks WHERE study_place_id = ? AND lesson_id = ? AND id = ? AND teacher_id = ? 
`,
		gocql.UUID(studyPlaceId),
		gocql.UUID(lessonId),
		gocql.UUID(markId),
		gocql.UUID(teacherId),
	).WithContext(ctx).Exec()
}

func (r *repository) EditMark(ctx context.Context, mark entities.Mark) error {
	return r.database.Query(`
UPDATE journal.marks SET  
	mark = ?,
	student_id = ?,
	updated_at = dateOf(now())
WHERE 
    study_place_id = ? AND lesson_id = ? AND id = ? AND teacher_id = ? 
IF EXISTS
`,
		mark.Mark,
		gocql.UUID(mark.StudentId),
		gocql.UUID(mark.StudyPlaceId),
		gocql.UUID(mark.LessonId),
		gocql.UUID(mark.ID),
		gocql.UUID(mark.TeacherId),
	).WithContext(ctx).Exec()
}
