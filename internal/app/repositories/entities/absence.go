package entities

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type Absence struct {
	entities.Timed
	ID           uuid.UUID
	StudyPlaceId uuid.UUID
	Absence      int
	StudentId    uuid.UUID
	TeacherId    uuid.UUID
	LessonId     uuid.UUID
}
