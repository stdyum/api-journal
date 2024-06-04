package controllers

import (
	"context"

	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-journal/internal/app/controllers/dto"
	. "github.com/stdyum/api-journal/internal/app/controllers/models"
	"github.com/stdyum/api-journal/internal/modules/schedule"
	"github.com/stdyum/api-journal/internal/modules/types_registry"
)

func (c *controller) GetOptions(ctx context.Context, enrollment models.Enrollment) ([]dto.OptionResponse, error) {
	optionsBuilder := NewOptionsBuilder()
	if enrollment.Role == models.RoleStudent {
		options, err := c.getStudentOptions(ctx, enrollment)
		if err != nil {
			return nil, err
		}

		optionsBuilder.Append(options...)
	}

	hasPermissionToViewAll := enrollment.Permissions.Assert("viewJournal") == nil
	if enrollment.Role == models.RoleTeacher && !hasPermissionToViewAll {
		options, err := c.getTeacherOptions(ctx, enrollment)
		if err != nil {
			return nil, err
		}

		optionsBuilder.Append(options...)
	}

	if hasPermissionToViewAll {
		options, err := c.getAllOptions(ctx, enrollment)
		if err != nil {
			return nil, err
		}

		optionsBuilder.Append(options...)
	}

	options, typesIds := optionsBuilder.Build()
	typesIds.Token = enrollment.Token
	typesIds.StudyPlaceId = enrollment.StudyPlaceId

	typesModels, err := c.typesRegistry.GetTypesById(ctx, typesIds)
	if err != nil {
		return nil, err
	}

	for i := range options {
		options[i].Subject = typesModels.Subjects[options[i].Subject.ID]
		options[i].Group = typesModels.Groups[options[i].Group.ID]
		options[i].Teacher = typesModels.Teachers[options[i].Teacher.ID]
	}

	return uslices.MapFunc(options, func(item Option) dto.OptionResponse {
		return dto.OptionResponse{
			Type: item.Type,
			Subject: dto.OptionResponseSubject{
				ID:   item.Subject.ID,
				Name: item.Subject.Name,
			},
			Group: dto.OptionResponseGroup{
				ID:   item.Group.ID,
				Name: item.Group.Name,
			},
			Teacher: dto.OptionResponseTeacher{
				ID:   item.Teacher.ID,
				Name: item.Teacher.Name,
			},
		}
	}), nil
}

func (c *controller) getStudentOptions(ctx context.Context, enrollment models.Enrollment) ([]Option, error) {
	groups, err := c.typesRegistry.GetStudentGroups(ctx, types_registry.GetStudentGroupsRequest{
		Token:        enrollment.Token,
		StudyPlaceId: enrollment.StudyPlaceId,
		StudentId:    enrollment.TypeId,
	})
	if err != nil {
		return nil, err
	}

	return uslices.MapFunc(groups, func(item models.Group) Option {
		return Option{
			Type: "student",
			Group: models.Group{
				ID: item.ID,
			},
		}
	}), nil
}

func (c *controller) getTeacherOptions(ctx context.Context, enrollment models.Enrollment) ([]Option, error) {
	return c.getOptionsAccordingToScheduleEntitiesFilter(ctx, enrollment, schedule.EntriesFilter{TeacherId: enrollment.TypeId})
}

func (c *controller) getAllOptions(ctx context.Context, enrollment models.Enrollment) ([]Option, error) {
	return c.getOptionsAccordingToScheduleEntitiesFilter(ctx, enrollment, schedule.EntriesFilter{})
}

func (c *controller) getOptionsAccordingToScheduleEntitiesFilter(ctx context.Context, enrollment models.Enrollment, entriesFilter schedule.EntriesFilter) ([]Option, error) {
	entriesFilter.Token = enrollment.Token
	entriesFilter.StudyPlaceId = enrollment.StudyPlaceId

	entries, err := c.schedule.GetUniqueEntries(ctx, entriesFilter)
	if err != nil {
		return nil, err
	}

	options := make([]Option, len(entries))
	for i, entry := range entries {
		options[i] = Option{
			Type: "group",
			Subject: models.Subject{
				ID: entry.SubjectId,
			},
			Group: models.Group{
				ID: entry.GroupId,
			},
			Teacher: models.Teacher{
				ID: entry.TeacherId,
			},
		}
	}

	return options, nil
}
