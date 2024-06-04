package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-journal/internal/app/controllers/dto"
)

func (h *http) GetJournal(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	tp := ctx.Query("type")

	subjectId, err := ctx.QueryUUID("subjectId")
	if err != nil && tp == "group" {
		_ = ctx.Error(err)
		return
	}

	groupId, err := ctx.QueryUUID("groupId")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	teacherId, err := ctx.QueryUUID("teacherId")
	if err != nil && tp == "group" {
		_ = ctx.Error(err)
		return
	}

	req := dto.GetJournalRequest{
		Type:      tp,
		SubjectId: subjectId,
		GroupId:   groupId,
		TeacherId: teacherId,
	}

	journal, err := h.controller.GetJournal(ctx, enrollment, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, journal)
}

func (h *http) GetOptions(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	options, err := h.controller.GetOptions(ctx, enrollment)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, options)
}

func (h *http) AddMark(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var req dto.AddMarkRequest
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	mark, err := h.controller.AddMark(ctx, enrollment, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, mark)
}

func (h *http) DeleteMark(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	id, err := ctx.UUIDParam("id")
	if err != nil {
		return
	}

	lessonId, err := ctx.QueryUUID("lessonId")
	if err != nil {
		return
	}

	req := dto.DeleteMarkRequest{
		Id:       id,
		LessonId: lessonId,
	}
	err = h.controller.DeleteMark(ctx, enrollment, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) EditMark(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var req dto.EditMarkRequest
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	err := h.controller.EditMark(ctx, enrollment, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) GetLessonInfo(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	lessonId, err := ctx.UUIDParam("id")
	if err != nil {
		return
	}

	req := dto.GetLessonInfoRequest{
		LessonId: lessonId,
	}
	lessonInfo, err := h.controller.GetLessonInfo(ctx, enrollment, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, lessonInfo)
}

func (h *http) AddLessonInfo(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var req dto.AddLessonInfoRequest
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	lessonInfo, err := h.controller.AddLessonInfo(ctx, enrollment, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, lessonInfo)
}

func (h *http) DeleteLessonInfo(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	id, err := ctx.UUIDParam("id")
	if err != nil {
		return
	}

	lessonId, err := ctx.QueryUUID("lessonId")
	if err != nil {
		return
	}

	req := dto.DeleteLessonInfoRequest{
		Id:       id,
		LessonId: lessonId,
	}
	err = h.controller.DeleteLessonInfo(ctx, enrollment, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) EditLessonInfo(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var req dto.EditLessonInfoRequest
	if err := ctx.BindJSON(&req); err != nil {
		return
	}

	err := h.controller.EditLessonInfo(ctx, enrollment, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}
