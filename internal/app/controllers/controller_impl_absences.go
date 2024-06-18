package controllers

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-journal/internal/app/controllers/dto"
	"github.com/stdyum/api-journal/internal/app/repositories/entities"
	"github.com/stdyum/api-journal/internal/modules/schedule"
)

func (c *controller) AddAbsences(ctx context.Context, enrollment models.Enrollment, request dto.AddAbsenceRequest) (dto.AddAbsenceResponse, error) {
	if enrollment.Role != models.RoleTeacher {
		return dto.AddAbsenceResponse{}, errors.New("role: no permission")
	}

	lesson, err := c.schedule.GetLessonById(ctx, schedule.GetLessonByIdRequest{
		Token:        enrollment.Token,
		StudyPlaceId: enrollment.StudyPlaceId,
		UUID:         request.LessonId,
	})
	if err != nil {
		return dto.AddAbsenceResponse{}, err
	}

	if lesson.TeacherId != enrollment.TypeId {
		return dto.AddAbsenceResponse{}, errors.New("teacherId: no permission")
	}

	absence := entities.Absence{
		ID:           uuid.New(),
		StudyPlaceId: enrollment.StudyPlaceId,
		Absence:      request.Absence,
		StudentId:    request.StudentId,
		TeacherId:    enrollment.TypeId,
		LessonId:     lesson.ID,
	}

	err = c.repository.AddAbsence(ctx, absence)
	if err != nil {
		return dto.AddAbsenceResponse{}, err
	}

	return dto.AddAbsenceResponse{
		ID:           absence.ID,
		StudyPlaceId: absence.StudyPlaceId,
		Absence:      absence.Absence,
		StudentId:    absence.StudentId,
		TeacherId:    absence.TeacherId,
		LessonId:     absence.LessonId,
	}, nil
}

func (c *controller) DeleteAbsence(ctx context.Context, enrollment models.Enrollment, request dto.DeleteAbsenceRequest) error {
	if enrollment.Role != models.RoleTeacher {
		return errors.New("role: no permission")
	}

	return c.repository.DeleteAbsence(ctx, enrollment.StudyPlaceId, request.LessonId, enrollment.TypeId, request.Id)
}

func (c *controller) EditAbsence(ctx context.Context, enrollment models.Enrollment, request dto.EditAbsenceRequest) error {
	if enrollment.Role != models.RoleTeacher {
		return errors.New("role: no permission")
	}

	absence := entities.Absence{
		ID:           request.Id,
		StudyPlaceId: enrollment.StudyPlaceId,
		Absence:      request.Absence,
		StudentId:    request.StudentId,
		TeacherId:    enrollment.TypeId,
		LessonId:     request.LessonId,
	}

	return c.repository.EditAbsence(ctx, absence)
}
