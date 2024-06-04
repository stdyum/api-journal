package modules

import (
	scheduleProto "github.com/stdyum/api-common/proto/impl/schedule"
	typesRegistryProto "github.com/stdyum/api-common/proto/impl/types_registry"

	"github.com/stdyum/api-journal/internal/modules/schedule"
	"github.com/stdyum/api-journal/internal/modules/types_registry"
)

type Modules struct {
	Schedule      schedule.Schedule
	TypesRegistry types_registry.TypesRegistry
}

func New(scheduleClient scheduleProto.ScheduleClient, typesRegistryClient typesRegistryProto.TypesRegistryClient) Modules {
	return Modules{
		Schedule:      schedule.NewSchedule(scheduleClient),
		TypesRegistry: types_registry.NewTypesRegistry(typesRegistryClient),
	}
}
