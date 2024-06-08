package controllers

import (
	"context"

	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-journal/internal/app/controllers/dto"
	"github.com/stdyum/api-journal/internal/app/repositories"
	"github.com/stdyum/api-journal/internal/modules/schedule"
	"github.com/stdyum/api-journal/internal/modules/types_registry"
)

type Controller interface {
	GetJournal(ctx context.Context, enrollment models.Enrollment, request dto.GetJournalRequest) (dto.JournalResponse, error)
	GetOptions(ctx context.Context, enrollment models.Enrollment, request dto.GetOptionsRequest) (dto.OptionsResponse, error)

	AddMark(ctx context.Context, enrollment models.Enrollment, request dto.AddMarkRequest) (dto.AddMarkResponse, error)
	DeleteMark(ctx context.Context, enrollment models.Enrollment, request dto.DeleteMarkRequest) error
	EditMark(ctx context.Context, enrollment models.Enrollment, request dto.EditMarkRequest) error

	GetLessonInfo(ctx context.Context, enrollment models.Enrollment, request dto.GetLessonInfoRequest) (dto.GetLessonInfoResponse, error)
	AddLessonInfo(ctx context.Context, enrollment models.Enrollment, request dto.AddLessonInfoRequest) (dto.AddLessonInfoResponse, error)
	DeleteLessonInfo(ctx context.Context, enrollment models.Enrollment, request dto.DeleteLessonInfoRequest) error
	EditLessonInfo(ctx context.Context, enrollment models.Enrollment, request dto.EditLessonInfoRequest) error
}

type controller struct {
	schedule      schedule.Schedule
	typesRegistry types_registry.TypesRegistry

	repository repositories.Repository
}

func New(repository repositories.Repository, schedule schedule.Schedule, typesRegistry types_registry.TypesRegistry) Controller {
	return &controller{
		schedule:      schedule,
		typesRegistry: typesRegistry,

		repository: repository,
	}
}
