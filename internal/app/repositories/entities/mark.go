package entities

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type Mark struct {
	entities.Timed
	ID           uuid.UUID
	StudyPlaceId uuid.UUID
	Mark         string
	StudentId    uuid.UUID
	TeacherId    uuid.UUID
	LessonId     uuid.UUID
}
