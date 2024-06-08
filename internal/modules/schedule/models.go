package schedule

import (
	"time"

	"github.com/google/uuid"
)

type GetLessonByIdRequest struct {
	Token        string
	StudyPlaceId uuid.UUID
	UUID         uuid.UUID
}

type Lesson struct {
	ID             uuid.UUID
	StudyPlaceId   uuid.UUID
	GroupId        uuid.UUID
	RoomId         uuid.UUID
	SubjectId      uuid.UUID
	TeacherId      uuid.UUID
	StartTime      time.Time
	EndTime        time.Time
	LessonIndex    int
	PrimaryColor   string
	SecondaryColor string
}

type EntriesFilter struct {
	Token        string
	StudyPlaceId uuid.UUID
	TeacherId    uuid.UUID
	GroupId      uuid.UUID
	SubjectId    uuid.UUID
	Cursor       string
	Limit        int
}

type Entries struct {
	List  []Entry
	Next  string
	Limit int
}

type Entry struct {
	TeacherId uuid.UUID
	GroupId   uuid.UUID
	SubjectId uuid.UUID
}
