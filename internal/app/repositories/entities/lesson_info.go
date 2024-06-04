package entities

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type LessonInfo struct {
	entities.Timed
	ID           uuid.UUID
	StudyPlaceId uuid.UUID
	LessonId     uuid.UUID
	TeacherId    uuid.UUID
	Title        string
	Description  string
	Homework     string
	Type         string
}
