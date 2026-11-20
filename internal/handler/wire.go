package handler

import (
	"github.com/google/wire"
	"go_load/internal/handler/grpc"
	"go_load/internal/handler/http"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
)