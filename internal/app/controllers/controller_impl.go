package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-journal/internal/app/controllers/dto"
	. "github.com/stdyum/api-journal/internal/app/controllers/models"
	"github.com/stdyum/api-journal/internal/modules/schedule"
	"github.com/stdyum/api-journal/internal/modules/types_registry"
)

func (c *controller) GetOptions(ctx context.Context, enrollment models.Enrollment, request dto.GetOptionsRequest) (dto.OptionsResponse, error) {
	if request.Limit == 0 {
		request.Limit = 10
	}

	optionsBuilder := NewOptionsBuilder()
	if enrollment.Role == models.RoleStudent && request.Cursor == "" {
		options, err := c.getStudentOptions(ctx, enrollment)
		if err != nil {
			return dto.OptionsResponse{}, err
		}

		optionsBuilder.Append(options...)
	}

	hasPermissionToViewAll := enrollment.Permissions.Assert("viewJournal") == nil
	if enrollment.Role == models.RoleTeacher && !hasPermissionToViewAll {
		groupIds, err := c.typesRegistry.GetGroupIdsWithStudents(ctx, types_registry.GetGroupIdsWithStudentsRequest{
			Token:        enrollment.Token,
			StudyPlaceId: enrollment.StudyPlaceId,
		})
		if err != nil {
			return dto.OptionsResponse{}, err
		}

		options, err := c.getTeacherOptions(ctx, enrollment, request, groupIds)
		if err != nil {
			return dto.OptionsResponse{}, err
		}

		optionsBuilder.AppendWithPagination(options)
	}

	if hasPermissionToViewAll {
		groupIds, err := c.typesRegistry.GetGroupIdsWithStudents(ctx, types_registry.GetGroupIdsWithStudentsRequest{
			Token:        enrollment.Token,
			StudyPlaceId: enrollment.StudyPlaceId,
		})
		if err != nil {
			return dto.OptionsResponse{}, err
		}

		options, err := c.getAllOptions(ctx, enrollment, request, groupIds)
		if err != nil {
			return dto.OptionsResponse{}, err
		}

		optionsBuilder.AppendWithPagination(options)
	}

	options, typesIds := optionsBuilder.Build()
	typesIds.Token = enrollment.Token
	typesIds.StudyPlaceId = enrollment.StudyPlaceId

	typesModels, err := c.typesRegistry.GetTypesById(ctx, typesIds)
	if err != nil {
		return dto.OptionsResponse{}, err
	}

	for i := range options.Options {
		options.Options[i].Subject = typesModels.Subjects[options.Options[i].Subject.ID]
		options.Options[i].Group = typesModels.Groups[options.Options[i].Group.ID]
		options.Options[i].Teacher = typesModels.Teachers[options.Options[i].Teacher.ID]
	}

	return dto.OptionsResponse{
		Options: uslices.MapFunc(options.Options, func(item Option) dto.OptionResponse {
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
				Editable: item.Editable,
			}
		}),
		Next:  options.Next,
		Limit: options.Limit,
	}, nil
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
			Editable: false,
		}
	}), nil
}

func (c *controller) getTeacherOptions(ctx context.Context, enrollment models.Enrollment, request dto.GetOptionsRequest, ids []uuid.UUID) (OptionsWithPagination, error) {
	return c.getOptionsAccordingToScheduleEntitiesFilter(ctx, enrollment, schedule.EntriesFilter{
		TeacherId: enrollment.TypeId,
		Cursor:    request.Cursor,
		Limit:     request.Limit,
		GroupIds:  ids,
	})
}

func (c *controller) getAllOptions(ctx context.Context, enrollment models.Enrollment, request dto.GetOptionsRequest, ids []uuid.UUID) (OptionsWithPagination, error) {
	return c.getOptionsAccordingToScheduleEntitiesFilter(ctx, enrollment, schedule.EntriesFilter{
		Cursor:   request.Cursor,
		Limit:    request.Limit,
		GroupIds: ids,
	})
}

func (c *controller) getOptionsAccordingToScheduleEntitiesFilter(ctx context.Context, enrollment models.Enrollment, entriesFilter schedule.EntriesFilter) (OptionsWithPagination, error) {
	entriesFilter.Token = enrollment.Token
	entriesFilter.StudyPlaceId = enrollment.StudyPlaceId

	entries, err := c.schedule.GetUniqueEntries(ctx, entriesFilter)
	if err != nil {
		return OptionsWithPagination{}, err
	}

	options := make([]Option, len(entries.List))
	for i, entry := range entries.List {
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
			Editable: enrollment.Role == models.RoleTeacher && enrollment.TypeId == entry.TeacherId,
		}
	}

	return OptionsWithPagination{
		Options: options,
		Next:    entries.Next,
		Limit:   entries.Limit,
	}, nil
}
