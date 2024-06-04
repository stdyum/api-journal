package handlers

import (
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/proto/impl/studyplaces"
	"github.com/stdyum/api-journal/internal/app/controllers"
)

type GRPC interface {
	grpc.Routes
	studyplaces.StudyplacesServer
}

type gRPC struct {
	studyplaces.UnimplementedStudyplacesServer

	controller controllers.Controller
}

func NewGRPC(controller controllers.Controller) GRPC {
	return &gRPC{
		controller: controller,
	}
}
