package types_registry

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	proto "github.com/stdyum/api-common/proto/impl/types_registry"
	"github.com/stdyum/api-common/uslices"
)

type TypesRegistry interface {
	GetTypesById(ctx context.Context, typesIds TypesIds) (TypesModels, error)
	GetStudentsInGroup(ctx context.Context, req GetStudentsRequest) ([]models.Student, error)
	GetStudentGroups(ctx context.Context, req GetStudentGroupsRequest) ([]models.Group, error)
}

type typesRegistry struct {
	server proto.TypesRegistryClient
}

func NewTypesRegistry(server proto.TypesRegistryClient) TypesRegistry {
	return &typesRegistry{
		server: server,
	}
}

func (t *typesRegistry) GetTypesById(ctx context.Context, typesIds TypesIds) (TypesModels, error) {
	in := proto.TypesIds{
		Token:        typesIds.Token,
		StudyPlaceId: typesIds.StudyPlaceId.String(),
		GroupsIds:    uslices.MapFunc(typesIds.GroupsIds, func(item uuid.UUID) string { return item.String() }),
		RoomsIds:     uslices.MapFunc(typesIds.RoomsIds, func(item uuid.UUID) string { return item.String() }),
		StudentIds:   uslices.MapFunc(typesIds.StudentIds, func(item uuid.UUID) string { return item.String() }),
		SubjectsIds:  uslices.MapFunc(typesIds.SubjectsIds, func(item uuid.UUID) string { return item.String() }),
		TeachersIds:  uslices.MapFunc(typesIds.TeachersIds, func(item uuid.UUID) string { return item.String() }),
	}

	types, err := t.server.GetTypesByIds(ctx, &in)
	if err != nil {
		return TypesModels{}, err
	}

	groups, err := t.mapGroups(types.Groups)
	if err != nil {
		return TypesModels{}, err
	}

	rooms, err := t.mapRooms(types.Rooms)
	if err != nil {
		return TypesModels{}, err
	}

	students, err := t.mapStudents(types.Students)
	if err != nil {
		return TypesModels{}, err
	}

	subjects, err := t.mapSubjects(types.Subjects)
	if err != nil {
		return TypesModels{}, err
	}

	teachers, err := t.mapTeachers(types.Teachers)
	if err != nil {
		return TypesModels{}, err
	}

	return TypesModels{
		Groups:   groups,
		Rooms:    rooms,
		Student:  students,
		Subjects: subjects,
		Teachers: teachers,
	}, nil
}

func (t *typesRegistry) GetStudentsInGroup(ctx context.Context, req GetStudentsRequest) ([]models.Student, error) {
	in := proto.GroupId{
		Token:        req.Token,
		StudyPlaceId: req.StudyPlaceId.String(),
		Uuid:         req.GroupId.String(),
	}
	students, err := t.server.GetStudentsInGroup(ctx, &in)
	if err != nil {
		return nil, err
	}

	return uslices.MapFuncErr(students.List, func(item *proto.Student) (models.Student, error) {
		id, err := uuid.Parse(item.Id)
		if err != nil {
			return models.Student{}, err
		}

		return models.Student{
			ID:   id,
			Name: item.Name,
		}, nil
	})
}

func (t *typesRegistry) GetStudentGroups(ctx context.Context, req GetStudentGroupsRequest) ([]models.Group, error) {
	in := proto.StudentId{
		Token:        req.Token,
		StudyPlaceId: req.StudyPlaceId.String(),
		Uuid:         req.StudentId.String(),
	}
	groups, err := t.server.GetStudentGroups(ctx, &in)
	if err != nil {
		return nil, err
	}

	return uslices.MapFuncErr(groups.List, func(item *proto.Group) (models.Group, error) {
		id, err := uuid.Parse(item.Id)
		if err != nil {
			return models.Group{}, err
		}

		return models.Group{
			ID:   id,
			Name: item.Name,
		}, nil
	})
}

func (t *typesRegistry) mapGroups(groups map[string]*proto.Group) (map[uuid.UUID]models.Group, error) {
	return uslices.MapMapFuncErr(groups, func(key string, item *proto.Group) (uuid.UUID, models.Group, error) {
		id, err := uuid.Parse(key)
		if err != nil {
			return uuid.Nil, models.Group{}, err
		}

		return id, models.Group{
			ID:   id,
			Name: item.Name,
		}, nil
	})
}

func (t *typesRegistry) mapRooms(rooms map[string]*proto.Room) (map[uuid.UUID]models.Room, error) {
	return uslices.MapMapFuncErr(rooms, func(key string, item *proto.Room) (uuid.UUID, models.Room, error) {
		id, err := uuid.Parse(key)
		if err != nil {
			return uuid.Nil, models.Room{}, err
		}

		return id, models.Room{
			ID:   id,
			Name: item.Name,
		}, nil
	})
}

func (t *typesRegistry) mapStudents(students map[string]*proto.Student) (map[uuid.UUID]models.Student, error) {
	return uslices.MapMapFuncErr(students, func(key string, item *proto.Student) (uuid.UUID, models.Student, error) {
		id, err := uuid.Parse(key)
		if err != nil {
			return uuid.Nil, models.Student{}, err
		}

		return id, models.Student{
			ID:   id,
			Name: item.Name,
		}, nil
	})
}

func (t *typesRegistry) mapSubjects(subjects map[string]*proto.Subject) (map[uuid.UUID]models.Subject, error) {
	return uslices.MapMapFuncErr(subjects, func(key string, item *proto.Subject) (uuid.UUID, models.Subject, error) {
		id, err := uuid.Parse(key)
		if err != nil {
			return uuid.Nil, models.Subject{}, err
		}

		return id, models.Subject{
			ID:   id,
			Name: item.Name,
		}, nil
	})
}

func (t *typesRegistry) mapTeachers(teachers map[string]*proto.Teacher) (map[uuid.UUID]models.Teacher, error) {
	return uslices.MapMapFuncErr(teachers, func(key string, item *proto.Teacher) (uuid.UUID, models.Teacher, error) {
		id, err := uuid.Parse(key)
		if err != nil {
			return uuid.Nil, models.Teacher{}, err
		}

		return id, models.Teacher{
			ID:   id,
			Name: item.Name,
		}, nil
	})
}
