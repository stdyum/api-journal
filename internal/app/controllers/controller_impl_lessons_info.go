package controllers

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-journal/internal/app/controllers/dto"
	"github.com/stdyum/api-journal/internal/app/repositories/entities"
)

func (c *controller) GetLessonInfo(ctx context.Context, enrollment models.Enrollment, request dto.GetLessonInfoRequest) (dto.GetLessonInfoResponse, error) {
	lessonInfo, err := c.repository.GetLessonInfoByLessonId(ctx, enrollment.StudyPlaceId, request.LessonId)
	if err != nil {
		return dto.GetLessonInfoResponse{}, err
	}

	return dto.GetLessonInfoResponse{
		ID:           lessonInfo.ID,
		StudyPlaceId: lessonInfo.StudyPlaceId,
		LessonId:     lessonInfo.LessonId,
		TeacherId:    lessonInfo.TeacherId,
		Title:        lessonInfo.Title,
		Description:  lessonInfo.Description,
		Homework:     lessonInfo.Homework,
		Type:         lessonInfo.Type,
	}, nil
}

func (c *controller) AddLessonInfo(ctx context.Context, enrollment models.Enrollment, request dto.AddLessonInfoRequest) (dto.AddLessonInfoResponse, error) {
	if enrollment.Role != models.RoleTeacher {
		return dto.AddLessonInfoResponse{}, errors.New("role: no permission")
	}

	// TODO
	//lesson, err := c.schedule.GetLessonById(ctx, schedule.GetLessonByIdRequest{
	//	Token:        enrollment.Token,
	//	StudyPlaceId: enrollment.StudyPlaceId,
	//	UUID:         request.LessonId,
	//})
	//if err != nil {
	//	return dto.AddLessonInfoResponse{}, err
	//}
	//
	//if lesson.TeacherId != enrollment.TypeId {
	//	return dto.AddLessonInfoResponse{}, errors.New("teacherId: no permission")
	//}

	lessonInfo := entities.LessonInfo{
		ID:           uuid.New(),
		StudyPlaceId: enrollment.StudyPlaceId,
		LessonId:     request.LessonId,
		TeacherId:    enrollment.TypeId,
		Title:        request.Title,
		Description:  request.Description,
		Homework:     request.Homework,
		Type:         request.Type,
	}

	err := c.repository.AddLessonInfo(ctx, lessonInfo)
	if err != nil {
		return dto.AddLessonInfoResponse{}, err
	}

	return dto.AddLessonInfoResponse{
		ID:           lessonInfo.ID,
		StudyPlaceId: lessonInfo.StudyPlaceId,
		LessonId:     lessonInfo.LessonId,
		TeacherId:    lessonInfo.TeacherId,
		Title:        lessonInfo.Title,
		Description:  lessonInfo.Description,
		Homework:     lessonInfo.Homework,
		Type:         lessonInfo.Type,
	}, nil
}

func (c *controller) DeleteLessonInfo(ctx context.Context, enrollment models.Enrollment, request dto.DeleteLessonInfoRequest) error {
	if enrollment.Role != models.RoleTeacher {
		return errors.New("role: no permission")
	}

	return c.repository.DeleteLessonInfo(ctx, enrollment.StudyPlaceId, request.LessonId, enrollment.TypeId, request.Id)
}

func (c *controller) EditLessonInfo(ctx context.Context, enrollment models.Enrollment, request dto.EditLessonInfoRequest) error {
	if enrollment.Role != models.RoleTeacher {
		return errors.New("role: no permission")
	}

	lessonInfo := entities.LessonInfo{
		ID:           request.Id,
		StudyPlaceId: enrollment.StudyPlaceId,
		LessonId:     request.LessonId,
		TeacherId:    enrollment.TypeId,
		Title:        request.Title,
		Description:  request.Description,
		Homework:     request.Homework,
		Type:         request.Type,
	}

	return c.repository.EditLessonInfo(ctx, lessonInfo)
}
