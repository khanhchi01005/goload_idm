package dataacess

import(
	"go_load/internal/dataacess/database"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	database.WireSet,
)