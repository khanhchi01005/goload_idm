//go:build wireinject

package wiring

import (
	"go_load/internal/configs"
	"go_load/internal/dataacess"
	"go_load/internal/handler"
	"go_load/internal/handler/grpc"
	"go_load/internal/logic"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	configs.WireSet,
	dataacess.WireSet,
	logic.WireSet,
	handler.WireSet,
)

func InitializeGRPCServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}
