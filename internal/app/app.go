package app

import (
	"github.com/gocql/gocql"
	"github.com/stdyum/api-common/proto/impl/schedule"
	"github.com/stdyum/api-common/proto/impl/types_registry"
	"github.com/stdyum/api-common/server"
	"github.com/stdyum/api-journal/internal/app/controllers"
	"github.com/stdyum/api-journal/internal/app/controllers/errors"
	"github.com/stdyum/api-journal/internal/app/handlers"
	"github.com/stdyum/api-journal/internal/app/repositories"
	"github.com/stdyum/api-journal/internal/modules"
)

func New(database *gocql.Session, scheduleClient schedule.ScheduleClient, typesRegistryClient types_registry.TypesRegistryClient) (server.Routes, error) {
	repo := repositories.New(database)

	mdl := modules.New(scheduleClient, typesRegistryClient)

	ctrl := controllers.New(repo, mdl.Schedule, mdl.TypesRegistry)

	errors.Register()

	httpHndl := handlers.NewHTTP(ctrl)
	grpcHndl := handlers.NewGRPC(ctrl)

	routes := server.Routes{
		GRPC: grpcHndl,
		HTTP: httpHndl,
	}

	return routes, nil
}
