package schedulers

import (
	"github.com/google/wire"
	"github.com/kainonly/cronx/api/jobs"
	"github.com/kainonly/cronx/api/storage"
	"github.com/kainonly/cronx/common"
)

var Provides = wire.NewSet(
	wire.Struct(new(Controller), "*"),
	wire.Struct(new(Service), "*"),
)

type Controller struct {
	V *common.Values

	SchedulersX *Service
}

type Service struct {
	*common.Inject

	StorageX *storage.Service
	JobsX    *jobs.Service
}

type M = map[string]any
