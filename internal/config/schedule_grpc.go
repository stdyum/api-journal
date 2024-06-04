package config

import (
	"github.com/stdyum/api-common/proto/impl/schedule"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ScheduleGRpcConfig struct {
	URL string `env:"URL"`
}

func ConnectToScheduleServer(config ScheduleGRpcConfig) (schedule.ScheduleClient, error) {
	conn, err := grpc.Dial(config.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return schedule.NewScheduleClient(conn), nil
}
