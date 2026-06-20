package storage

import (
	"github.com/google/wire"
	"github.com/kainonly/cronx/common"
)

var Provides = wire.NewSet(
	wire.Struct(new(Controller), "*"),
	wire.Struct(new(Service), "*"),
)

type Controller struct {
	V *common.Values

	StorageX *Service
}

type Service struct {
	*common.Inject
}

type M = map[string]any
