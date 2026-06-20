package api

import (
	"context"

	"github.com/kainonly/cronx/api/index"
	"github.com/kainonly/cronx/api/jobs"
	"github.com/kainonly/cronx/api/schedulers"
	"github.com/kainonly/cronx/api/storage"
	"github.com/kainonly/cronx/common"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/google/wire"
)

var Provides = wire.NewSet(
	index.Provides,
	jobs.Provides,
	schedulers.Provides,
	storage.Provides,
)

type API struct {
	*common.Inject

	Hertz      *server.Hertz
	Index      *index.Controller
	IndexX     *index.Service
	Jobs       *jobs.Controller
	Schedulers *schedulers.Controller
	Storage    *storage.Controller
}

func (x *API) Initialize(ctx context.Context) (h *server.Hertz, err error) {
	x.Hertz.GET("", x.Index.Ping)

	_jobs := x.Hertz.Group("jobs")
	{
		_jobs.POST(`set`, x.Jobs.Set)
		_jobs.POST(`remove`, x.Jobs.Remove)
	}
	_schedulers := x.Hertz.Group("schedulers")
	{
		_schedulers.POST(`set`, x.Schedulers.Set)
		_schedulers.POST(`start`, x.Schedulers.Start)
		_schedulers.POST(`stop`, x.Schedulers.Stop)
		_schedulers.POST(`remove`, x.Schedulers.Remove)
	}
	_storage := x.Hertz.Group("storage")
	{
		_storage.GET(``, x.Storage.List)
	}
	return x.Hertz, nil
}
