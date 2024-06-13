package schedule

import (
	"context"
	"time"

	"github.com/google/uuid"
	proto "github.com/stdyum/api-common/proto/impl/schedule"
	"github.com/stdyum/api-common/uslices"
)

type Schedule interface {
	GetLessonById(ctx context.Context, request GetLessonByIdRequest) (Lesson, error)
	GetLessons(ctx context.Context, filter EntriesFilter) ([]Lesson, error)
	GetUniqueEntries(ctx context.Context, filter EntriesFilter) (Entries, error)
}

type schedule struct {
	server proto.ScheduleClient
}

func NewSchedule(server proto.ScheduleClient) Schedule {
	return &schedule{
		server: server,
	}
}

func (s *schedule) GetLessonById(ctx context.Context, request GetLessonByIdRequest) (Lesson, error) {
	in := proto.UUID{
		Token:        request.Token,
		StudyPlaceId: request.StudyPlaceId.String(),
		Uuid:         request.UUID.String(),
	}

	out, err := s.server.GetLessonById(ctx, &in)
	if err != nil {
		return Lesson{}, err
	}

	id, err := uuid.Parse(out.Id)
	if err != nil {
		return Lesson{}, err
	}

	studyPlaceId, err := uuid.Parse(out.StudyPlaceId)
	if err != nil {
		return Lesson{}, err
	}

	groupId, err := uuid.Parse(out.GroupId)
	if err != nil {
		return Lesson{}, err
	}

	roomId, err := uuid.Parse(out.RoomId)
	if err != nil {
		return Lesson{}, err
	}

	subjectId, err := uuid.Parse(out.SubjectId)
	if err != nil {
		return Lesson{}, err
	}

	teacherId, err := uuid.Parse(out.TeacherId)
	if err != nil {
		return Lesson{}, err
	}

	return Lesson{
		ID:             id,
		StudyPlaceId:   studyPlaceId,
		GroupId:        groupId,
		RoomId:         roomId,
		SubjectId:      subjectId,
		TeacherId:      teacherId,
		StartTime:      time.Unix(out.StartTime, 0),
		EndTime:        time.Unix(out.EndTime, 0),
		LessonIndex:    int(out.LessonIndex),
		PrimaryColor:   out.PrimaryColor,
		SecondaryColor: out.SecondaryColor,
	}, nil
}

func (s *schedule) GetLessons(ctx context.Context, filter EntriesFilter) ([]Lesson, error) {
	in := proto.EntriesFilter{
		Token:        filter.Token,
		StudyPlaceId: filter.StudyPlaceId.String(),
		TeacherId:    filter.TeacherId.String(),
		GroupIds: uslices.MapFunc(filter.GroupIds, func(item uuid.UUID) string {
			return item.String()
		}),
		SubjectId: filter.SubjectId.String(),
	}

	out, err := s.server.GetLessons(ctx, &in)
	if err != nil {
		return nil, err
	}

	return uslices.MapFuncErr(out.List, func(item *proto.Lesson) (Lesson, error) {
		id, err := uuid.Parse(item.Id)
		if err != nil {
			return Lesson{}, err
		}

		studyPlaceId, err := uuid.Parse(item.StudyPlaceId)
		if err != nil {
			return Lesson{}, err
		}

		groupId, err := uuid.Parse(item.GroupId)
		if err != nil {
			return Lesson{}, err
		}

		roomId, err := uuid.Parse(item.RoomId)
		if err != nil {
			return Lesson{}, err
		}

		subjectId, err := uuid.Parse(item.SubjectId)
		if err != nil {
			return Lesson{}, err
		}

		teacherId, err := uuid.Parse(item.TeacherId)
		if err != nil {
			return Lesson{}, err
		}

		return Lesson{
			ID:             id,
			StudyPlaceId:   studyPlaceId,
			GroupId:        groupId,
			RoomId:         roomId,
			SubjectId:      subjectId,
			TeacherId:      teacherId,
			StartTime:      time.Unix(item.StartTime, 0),
			EndTime:        time.Unix(item.EndTime, 0),
			LessonIndex:    int(item.LessonIndex),
			PrimaryColor:   item.PrimaryColor,
			SecondaryColor: item.SecondaryColor,
		}, nil
	})
}

func (s *schedule) GetUniqueEntries(ctx context.Context, filter EntriesFilter) (Entries, error) {
	in := proto.EntriesFilter{
		Token:        filter.Token,
		StudyPlaceId: filter.StudyPlaceId.String(),
		TeacherId:    filter.TeacherId.String(),
		GroupIds: uslices.MapFunc(filter.GroupIds, func(item uuid.UUID) string {
			return item.String()
		}),
		SubjectId: filter.SubjectId.String(),
		Cursor:    filter.Cursor,
		Limit:     int32(filter.Limit),
	}

	out, err := s.server.GetUniqueEntries(ctx, &in)
	if err != nil {
		return Entries{}, err
	}

	entries := make([]Entry, len(out.List))
	for i, entry := range out.List {
		entries[i], err = s.convertGRpcEntryToModel(entry)
		if err != nil {
			return Entries{}, err
		}
	}

	return Entries{
		List:  entries,
		Next:  out.Next,
		Limit: int(out.Limit),
	}, nil
}

func (s *schedule) convertGRpcEntryToModel(in *proto.Entry) (Entry, error) {
	teacherId, err := uuid.Parse(in.TeacherId)
	if err != nil {
		return Entry{}, err
	}

	groupId, err := uuid.Parse(in.GroupId)
	if err != nil {
		return Entry{}, err
	}

	subjectId, err := uuid.Parse(in.SubjectId)
	if err != nil {
		return Entry{}, err
	}

	return Entry{
		TeacherId: teacherId,
		GroupId:   groupId,
		SubjectId: subjectId,
	}, nil
}
