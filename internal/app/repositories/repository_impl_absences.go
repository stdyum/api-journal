package repositories

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-journal/internal/app/repositories/entities"
)

func (r *repository) GetStudentAbsences(ctx context.Context, studyPlaceId uuid.UUID, lessonIds []uuid.UUID, studentId uuid.UUID) ([]entities.Absence, error) {
	scanner := r.database.Query(
		"SELECT id, study_place_id, absence, student_id, teacher_id, lesson_id FROM absences WHERE study_place_id = ? AND student_id = ? ALLOW FILTERING",
		gocql.UUID(studyPlaceId),
		//uslices.MapFunc(lessonIds, func(item uuid.UUID) gocql.UUID { return gocql.UUID(item) }),
		gocql.UUID(studentId),
	).
		WithContext(ctx).
		Iter().
		Scanner()

	return databases.ScanArray(scanner, r.scanAbsence)
}

func (r *repository) GetLessonsAbsences(ctx context.Context, studyPlaceId uuid.UUID, lessonIds []uuid.UUID) ([]entities.Absence, error) {
	scanner := r.database.Query(
		"SELECT id, study_place_id, absence, student_id, teacher_id, lesson_id FROM absences WHERE study_place_id = ? AND lesson_id IN ? ALLOW FILTERING",
		gocql.UUID(studyPlaceId),
		uslices.MapFunc(lessonIds, func(item uuid.UUID) gocql.UUID { return gocql.UUID(item) }),
	).
		WithContext(ctx).
		Iter().
		Scanner()

	return databases.ScanArray(scanner, r.scanAbsence)
}

func (r *repository) AddAbsence(ctx context.Context, absence entities.Absence) error {
	return r.database.Query(`
INSERT INTO journal.absences 
    (id, study_place_id, absence, student_id, teacher_id, lesson_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, dateOf(now()), dateOf(now()));`,
		gocql.UUID(absence.ID),
		gocql.UUID(absence.StudyPlaceId),
		absence.Absence,
		gocql.UUID(absence.StudentId),
		gocql.UUID(absence.TeacherId),
		gocql.UUID(absence.LessonId),
	).
		WithContext(ctx).
		Exec()
}

func (r *repository) DeleteAbsence(ctx context.Context, studyPlaceId uuid.UUID, lessonId uuid.UUID, teacherId uuid.UUID, absenceId uuid.UUID) error {
	return r.database.Query(`
DELETE FROM journal.absences WHERE study_place_id = ? AND lesson_id = ? AND id = ? AND teacher_id = ? 
`,
		gocql.UUID(studyPlaceId),
		gocql.UUID(lessonId),
		gocql.UUID(absenceId),
		gocql.UUID(teacherId),
	).WithContext(ctx).Exec()
}

func (r *repository) EditAbsence(ctx context.Context, absence entities.Absence) error {
	return r.database.Query(`
UPDATE journal.absences SET  
	absence = ?,
	student_id = ?,
	updated_at = dateOf(now())
WHERE 
    study_place_id = ? AND lesson_id = ? AND id = ? AND teacher_id = ? 
IF EXISTS
`,
		absence.Absence,
		gocql.UUID(absence.StudentId),
		gocql.UUID(absence.StudyPlaceId),
		gocql.UUID(absence.LessonId),
		gocql.UUID(absence.ID),
		gocql.UUID(absence.TeacherId),
	).WithContext(ctx).Exec()
}
