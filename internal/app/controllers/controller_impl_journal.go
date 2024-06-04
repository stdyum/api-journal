package controllers

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-journal/internal/app/controllers/dto"
	"github.com/stdyum/api-journal/internal/app/repositories/entities"
	"github.com/stdyum/api-journal/internal/modules/schedule"
	"github.com/stdyum/api-journal/internal/modules/types_registry"
)

func (c *controller) GetJournal(ctx context.Context, enrollment models.Enrollment, request dto.GetJournalRequest) (dto.JournalResponse, error) {
	switch request.Type {
	case "group":
		return c.getGroupJournal(ctx, enrollment, request)
	case "student":
		return c.getStudentJournal(ctx, enrollment, request)
	default:
		return dto.JournalResponse{}, errors.New("unknown journal type")
	}
}

func (c *controller) getGroupJournal(ctx context.Context, enrollment models.Enrollment, request dto.GetJournalRequest) (dto.JournalResponse, error) {
	if enrollment.Permissions.Assert("viewJournal") != nil && (enrollment.Role != models.RoleTeacher || enrollment.TypeId != request.TeacherId) {
		return dto.JournalResponse{}, errors.New("role: no permission")
	}

	filter := schedule.EntriesFilter{
		Token:        enrollment.Token,
		StudyPlaceId: enrollment.StudyPlaceId,
		TeacherId:    request.TeacherId,
		GroupId:      request.GroupId,
		SubjectId:    request.SubjectId,
	}

	lessons, err := c.schedule.GetLessons(ctx, filter)
	if err != nil {
		return dto.JournalResponse{}, err
	}

	lessonMap := make(map[uuid.UUID]schedule.Lesson, len(lessons))
	lessonIds := make([]uuid.UUID, len(lessons))
	for i, lesson := range lessons {
		lessonMap[lesson.ID] = lesson
		lessonIds[i] = lesson.ID
	}

	students, err := c.typesRegistry.GetStudentsInGroup(ctx, types_registry.GetStudentsRequest{
		Token:        enrollment.Token,
		StudyPlaceId: enrollment.StudyPlaceId,
		GroupId:      request.GroupId,
	})
	if err != nil {
		return dto.JournalResponse{}, err
	}

	studentMap := make(map[uuid.UUID]models.Student, len(lessons))
	for _, student := range students {
		studentMap[student.ID] = student
	}

	marks, err := c.repository.GetLessonsMarks(ctx, enrollment.StudyPlaceId, lessonIds)
	if err != nil {
		return dto.JournalResponse{}, err
	}

	groupedMarks := uslices.GroupBy(marks, func(item entities.Mark) string {
		return item.LessonId.String() + item.StudentId.String()
	})

	cells := uslices.MapFunc(groupedMarks, func(marks []entities.Mark) dto.JournalCellResponse {
		return dto.JournalCellResponse{
			Point: dto.JournalCellPointResponse{
				RowId:    marks[0].StudentId.String(),
				ColumnId: marks[0].LessonId.String(),
			},
			Marks: uslices.MapFunc(marks, func(item entities.Mark) dto.JournalCellMarkResponse {
				return dto.JournalCellMarkResponse{
					Id:        item.ID.String(),
					Mark:      item.Mark,
					LessonId:  item.LessonId.String(),
					StudentId: item.StudentId.String(),
				}
			}),
		}
	})

	rows := uslices.MapFunc(students, func(item models.Student) dto.JournalRowResponse {
		return dto.JournalRowResponse{
			Id:    item.ID.String(),
			Title: item.Name,
		}
	})

	columns := uslices.MapFunc(lessons, func(item schedule.Lesson) dto.JournalColumnResponse {
		return dto.JournalColumnResponse{
			Id:    item.ID.String(),
			Title: item.StartTime.Format("2006-01-02 15:04"),
		}
	})

	return dto.JournalResponse{
		Rows:    rows,
		Columns: columns,
		Cells:   cells,
		Info: dto.JournalInfoResponse{
			Type: request.Type,
		},
	}, nil
}

func (c *controller) getStudentJournal(ctx context.Context, enrollment models.Enrollment, request dto.GetJournalRequest) (dto.JournalResponse, error) {
	if enrollment.Role != models.RoleStudent {
		return dto.JournalResponse{}, errors.New("role: no permission")
	}

	filter := schedule.EntriesFilter{
		Token:        enrollment.Token,
		StudyPlaceId: enrollment.StudyPlaceId,
		GroupId:      request.GroupId,
	}

	lessons, err := c.schedule.GetLessons(ctx, filter)
	if err != nil {
		return dto.JournalResponse{}, err
	}

	lessonMap := make(map[uuid.UUID]schedule.Lesson, len(lessons))
	lessonIds := make([]uuid.UUID, len(lessons))
	subjectIds := make([]uuid.UUID, len(lessons))
	subjects := make([]models.Subject, len(lessons))
	for i, lesson := range lessons {
		lessonMap[lesson.ID] = lesson
		lessonIds[i] = lesson.ID
		subjectIds[i] = lesson.SubjectId
		subjects[i] = models.Subject{
			ID: lesson.SubjectId,
		}
	}

	typesModels, err := c.typesRegistry.GetTypesById(ctx, types_registry.TypesIds{
		Token:        enrollment.Token,
		StudyPlaceId: enrollment.StudyPlaceId,
		SubjectsIds:  subjectIds,
	})
	if err != nil {
		return dto.JournalResponse{}, err
	}

	for i := range subjects {
		subjects[i].Name = typesModels.Subjects[subjects[i].ID].Name
	}

	marks, err := c.repository.GetStudentMarks(ctx, enrollment.StudyPlaceId, lessonIds, enrollment.TypeId)
	if err != nil {
		return dto.JournalResponse{}, err
	}

	groupedMarks := uslices.GroupBy(marks, func(item entities.Mark) string {
		return item.LessonId.String() + item.StudentId.String()
	})

	cells := uslices.MapFunc(groupedMarks, func(marks []entities.Mark) dto.JournalCellResponse {
		return dto.JournalCellResponse{
			Point: dto.JournalCellPointResponse{
				RowId:    lessonMap[marks[0].LessonId].SubjectId.String(),
				ColumnId: marks[0].LessonId.String(),
			},
			Marks: uslices.MapFunc(marks, func(item entities.Mark) dto.JournalCellMarkResponse {
				return dto.JournalCellMarkResponse{
					Id:        item.ID.String(),
					Mark:      item.Mark,
					LessonId:  item.LessonId.String(),
					StudentId: item.StudentId.String(),
				}
			}),
		}
	})

	rows := uslices.MapFunc(subjects, func(item models.Subject) dto.JournalRowResponse {
		return dto.JournalRowResponse{
			Id:    item.ID.String(),
			Title: item.Name,
		}
	})

	columns := uslices.MapFunc(lessons, func(item schedule.Lesson) dto.JournalColumnResponse {
		return dto.JournalColumnResponse{
			Id:    item.ID.String(),
			Title: item.StartTime.Format("2006-01-02 15:04"),
		}
	})

	return dto.JournalResponse{
		Rows:    rows,
		Columns: columns,
		Cells:   cells,
		Info: dto.JournalInfoResponse{
			Type: request.Type,
		},
	}, nil
}
