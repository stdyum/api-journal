package handlers

import (
	"github.com/stdyum/api-common/hc"
	confHttp "github.com/stdyum/api-common/http"
	"github.com/stdyum/api-journal/internal/app/controllers"
)

type HTTP interface {
	confHttp.Routes

	GetJournal(ctx *hc.Context)
	GetOptions(ctx *hc.Context)

	AddMark(ctx *hc.Context)
	DeleteMark(ctx *hc.Context)
	EditMark(ctx *hc.Context)

	GetLessonInfo(ctx *hc.Context)
	AddLessonInfo(ctx *hc.Context)
	DeleteLessonInfo(ctx *hc.Context)
	EditLessonInfo(ctx *hc.Context)

	AddAbsence(ctx *hc.Context)
	DeleteAbsence(ctx *hc.Context)
	EditAbsence(ctx *hc.Context)
}

type http struct {
	controller controllers.Controller
}

func NewHTTP(controller controllers.Controller) HTTP {
	return &http{
		controller: controller,
	}
}
