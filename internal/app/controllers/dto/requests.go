package dto

import (
	"github.com/google/uuid"
)

type GetOptionsRequest struct {
	Cursor string `form:"cursor"`
	Limit  int    `form:"limit"`
}

type GetJournalRequest struct {
	Type      string    `form:"type"`
	SubjectId uuid.UUID `form:"subjectId"`
	GroupId   uuid.UUID `form:"groupId"`
	TeacherId uuid.UUID `form:"teacherId'"`
}

type AddMarkRequest struct {
	Mark      string    `json:"mark"`
	StudentId uuid.UUID `json:"studentId"`
	LessonId  uuid.UUID `json:"lessonId"`
}

type EditMarkRequest struct {
	Id        uuid.UUID `json:"id"`
	Mark      string    `json:"mark"`
	StudentId uuid.UUID `json:"studentId"`
	LessonId  uuid.UUID `json:"lessonId"`
}

type DeleteMarkRequest struct {
	Id       uuid.UUID `json:"id"`
	LessonId uuid.UUID `json:"lessonId"`
}

type GetLessonInfoRequest struct {
	LessonId uuid.UUID `json:"lessonId"`
}

type AddLessonInfoRequest struct {
	LessonId    uuid.UUID `json:"lessonId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Homework    string    `json:"homework"`
	Type        string    `json:"type"`
}

type EditLessonInfoRequest struct {
	Id          uuid.UUID `json:"id"`
	LessonId    uuid.UUID `json:"lessonId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Homework    string    `json:"homework"`
	Type        string    `json:"type"`
}

type DeleteLessonInfoRequest struct {
	Id       uuid.UUID `json:"id"`
	LessonId uuid.UUID `json:"lessonId"`
}

type AddAbsenceRequest struct {
	Absence   int       `json:"absence"`
	StudentId uuid.UUID `json:"studentId"`
	LessonId  uuid.UUID `json:"lessonId"`
}

type EditAbsenceRequest struct {
	Id        uuid.UUID `json:"id"`
	Absence   int       `json:"absence"`
	StudentId uuid.UUID `json:"studentId"`
	LessonId  uuid.UUID `json:"lessonId"`
}

type DeleteAbsenceRequest struct {
	Id       uuid.UUID `json:"id"`
	LessonId uuid.UUID `json:"lessonId"`
}
