package models

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-journal/internal/app/repositories/entities"
)

type CellEntry struct {
	Mark    *entities.Mark
	Absence *entities.Absence

	LessonId  uuid.UUID
	StudentId uuid.UUID
}
