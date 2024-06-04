package internal

import (
	"github.com/stdyum/api-common/grpc/clients"
	"github.com/stdyum/api-journal/internal/app"
	"github.com/stdyum/api-journal/internal/config"
)

func App() error {
	db, err := config.ConnectToDatabase(config.Config.Database)
	if err != nil {
		return err
	}

	studyPlacesServer, err := config.ConnectToStudyPlacesServer(config.Config.StudyPlacesGRpc)
	if err != nil {
		return err
	}
	clients.StudyPlacesGRpcClient = studyPlacesServer

	typesRegistryService, err := config.ConnectToSTypesRegistryServer(config.Config.TypesRegistryGRpc)
	if err != nil {
		return err
	}

	scheduleService, err := config.ConnectToScheduleServer(config.Config.ScheduleGRPC)
	if err != nil {
		return err
	}

	routes, err := app.New(db, scheduleService, typesRegistryService)
	if err != nil {
		return err
	}

	routes.Ports = config.Config.Ports
	return routes.Run()
}
