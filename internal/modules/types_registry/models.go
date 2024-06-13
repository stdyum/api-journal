package types_registry

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
)

type TypesIds struct {
	Token        string
	StudyPlaceId uuid.UUID
	GroupsIds    []uuid.UUID
	RoomsIds     []uuid.UUID
	StudentIds   []uuid.UUID
	SubjectsIds  []uuid.UUID
	TeachersIds  []uuid.UUID
}

type TypesModels struct {
	Groups   map[uuid.UUID]models.Group
	Rooms    map[uuid.UUID]models.Room
	Student  map[uuid.UUID]models.Student
	Subjects map[uuid.UUID]models.Subject
	Teachers map[uuid.UUID]models.Teacher
}

type GetStudentsRequest struct {
	Token        string
	StudyPlaceId uuid.UUID
	GroupId      uuid.UUID
}

type GetStudentGroupsRequest struct {
	Token        string
	StudyPlaceId uuid.UUID
	StudentId    uuid.UUID
}

type GetGroupIdsWithStudentsRequest struct {
	Token        string
	StudyPlaceId uuid.UUID
}
