package handlers

import (
	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-common/http/middlewares"
	"google.golang.org/grpc"
)

func (h *http) ConfigureRoutes() *hc.Engine {
	engine := hc.New()
	engine.Use(hc.Recovery())

	group := engine.Group("api/v1", hc.Logger(), middlewares.ErrorMiddleware())
	{
		withAuth := group.Group("", middlewares.EnrollmentAuthMiddleware())

		{
			withAuth.GET("journal", h.GetJournal)
			withAuth.GET("options", h.GetOptions)

			marks := withAuth.Group("marks")
			{
				marks.POST("", h.AddMark)
				marks.DELETE(":id", h.DeleteMark)
				marks.PUT("", h.EditMark)
			}

			lessonInfo := withAuth.Group("lessons/info")
			{
				lessonInfo.GET(":id", h.GetLessonInfo)
				lessonInfo.POST("", h.AddLessonInfo)
				lessonInfo.DELETE(":id", h.DeleteLessonInfo)
				lessonInfo.PUT("", h.EditLessonInfo)
			}
		}
	}

	return engine
}

func (h *gRPC) ConfigureRoutes() *grpc.Server {
	return nil
}
