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

func (c *controller) AddMark(ctx context.Context, enrollment models.Enrollment, request dto.AddMarkRequest) (dto.AddMarkResponse, error) {
	if enrollment.Role != models.RoleTeacher {
		return dto.AddMarkResponse{}, errors.New("role: no permission")
	}

	lesson, err := c.schedule.GetLessonById(ctx, schedule.GetLessonByIdRequest{
		Token:        enrollment.Token,
		StudyPlaceId: enrollment.StudyPlaceId,
		UUID:         request.LessonId,
	})
	if err != nil {
		return dto.AddMarkResponse{}, err
	}

	if lesson.TeacherId != enrollment.TypeId {
		return dto.AddMarkResponse{}, errors.New("teacherId: no permission")
	}

	mark := entities.Mark{
		ID:           uuid.New(),
		StudyPlaceId: enrollment.StudyPlaceId,
		Mark:         request.Mark,
		StudentId:    request.StudentId,
		TeacherId:    enrollment.TypeId,
		LessonId:     lesson.ID,
	}

	err = c.repository.AddMark(ctx, mark)
	if err != nil {
		return dto.AddMarkResponse{}, err
	}

	return dto.AddMarkResponse{
		ID:           mark.ID,
		StudyPlaceId: mark.StudyPlaceId,
		Mark:         mark.Mark,
		StudentId:    mark.StudentId,
		TeacherId:    mark.TeacherId,
		LessonId:     mark.LessonId,
	}, nil
}

func (c *controller) DeleteMark(ctx context.Context, enrollment models.Enrollment, request dto.DeleteMarkRequest) error {
	if enrollment.Role != models.RoleTeacher {
		return errors.New("role: no permission")
	}

	return c.repository.DeleteMark(ctx, enrollment.StudyPlaceId, request.LessonId, enrollment.TypeId, request.Id)
}

func (c *controller) EditMark(ctx context.Context, enrollment models.Enrollment, request dto.EditMarkRequest) error {
	if enrollment.Role != models.RoleTeacher {
		return errors.New("role: no permission")
	}

	mark := entities.Mark{
		ID:           request.Id,
		StudyPlaceId: enrollment.StudyPlaceId,
		Mark:         request.Mark,
		StudentId:    request.StudentId,
		TeacherId:    enrollment.TypeId,
		LessonId:     request.LessonId,
	}

	return c.repository.EditMark(ctx, mark)
}
