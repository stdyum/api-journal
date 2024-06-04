package config

import (
	"github.com/stdyum/api-common/env"
	"github.com/stdyum/api-common/server"
)

var Config Model

type Model struct {
	Ports             server.PortConfig       `env:"PORT"`
	Database          DatabaseConfig          `env:"DATABASE"`
	StudyPlacesGRpc   StudyPlacesGRpcConfig   `env:"GRPC_STUDY_PLACES"`
	TypesRegistryGRpc TypesRegistryGRpcConfig `env:"GRPC_TYPES_REGISTRY"`
	ScheduleGRPC      ScheduleGRpcConfig      `env:"GRPC_SCHEDULE"`
}

func init() {
	err := env.Fill(&Config)
	if err != nil {
		panic("cannot fill config: " + err.Error())
	}
}
