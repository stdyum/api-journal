package dto

import (
	"github.com/google/uuid"
)

type OptionResponse struct {
	Type    string                `json:"type"`
	Subject OptionResponseSubject `json:"subject"`
	Group   OptionResponseGroup   `json:"group"`
	Teacher OptionResponseTeacher `json:"teacher"`
}

type OptionResponseSubject struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type OptionResponseGroup struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type OptionResponseTeacher struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type JournalResponse struct {
	Rows    []JournalRowResponse    `json:"rows"`
	Columns []JournalColumnResponse `json:"columns"`
	Cells   []JournalCellResponse   `json:"cells"`
	Info    JournalInfoResponse     `json:"info"`
}

type JournalInfoResponse struct {
	Type    string `json:"type"`
	Subject string `json:"subject"`
	Group   string `json:"group"`
	Teacher string `json:"teacher"`
}

type JournalRowResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type JournalColumnResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type JournalCellResponse struct {
	Point JournalCellPointResponse  `json:"point"`
	Marks []JournalCellMarkResponse `json:"marks"`
}

type JournalCellPointResponse struct {
	RowId    string `json:"rowId"`
	ColumnId string `json:"columnId"`
}

type JournalCellMarkResponse struct {
	Id        string `json:"id"`
	Mark      string `json:"mark"`
	LessonId  string `json:"lessonId"`
	StudentId string `json:"studentId"`
}

type AddMarkResponse struct {
	ID           uuid.UUID `json:"id"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	Mark         string    `json:"mark"`
	StudentId    uuid.UUID `json:"studentId"`
	TeacherId    uuid.UUID `json:"teacherId"`
	LessonId     uuid.UUID `json:"lessonId"`
}

type GetLessonInfoResponse struct {
	ID           uuid.UUID `json:"id"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	LessonId     uuid.UUID `json:"lessonId"`
	TeacherId    uuid.UUID `json:"teacherId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Homework     string    `json:"homework"`
	Type         string    `json:"type"`
}

type AddLessonInfoResponse struct {
	ID           uuid.UUID `json:"id"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	LessonId     uuid.UUID `json:"lessonId"`
	TeacherId    uuid.UUID `json:"teacherId"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Homework     string    `json:"homework"`
	Type         string    `json:"type"`
}
