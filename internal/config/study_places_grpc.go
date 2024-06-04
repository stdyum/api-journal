package config

import (
	"github.com/stdyum/api-common/proto/impl/studyplaces"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type StudyPlacesGRpcConfig struct {
	URL string `env:"URL"`
}

func ConnectToStudyPlacesServer(config StudyPlacesGRpcConfig) (studyplaces.StudyplacesClient, error) {
	conn, err := grpc.Dial(config.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return studyplaces.NewStudyplacesClient(conn), nil
}
