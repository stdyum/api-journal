package config

import (
	"github.com/stdyum/api-common/proto/impl/types_registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TypesRegistryGRpcConfig struct {
	URL string `env:"URL"`
}

func ConnectToSTypesRegistryServer(config TypesRegistryGRpcConfig) (types_registry.TypesRegistryClient, error) {
	conn, err := grpc.Dial(config.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return types_registry.NewTypesRegistryClient(conn), nil
}
