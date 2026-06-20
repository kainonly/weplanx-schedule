package api

import (
	"context"

	"github.com/kainonly/cronx/api/index"
	"github.com/kainonly/cronx/api/jobs"
	"github.com/kainonly/cronx/api/schedulers"
	"github.com/kainonly/cronx/common"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/google/wire"
)

var Provides = wire.NewSet(
	index.Provides,
	jobs.Provides,
	schedulers.Provides,
)

type API struct {
	*common.Inject

	Hertz      *server.Hertz
	Index      *index.Controller
	IndexX     *index.Service
	Jobs       *jobs.Controller
	Schedulers *schedulers.Controller
}

func (x *API) Initialize(ctx context.Context) (h *server.Hertz, err error) {
	x.Hertz.GET("", x.Index.Ping)

	_jobs := x.Hertz.Group("jobs")
	{
		_jobs.POST(`create`, x.Jobs.Create)
		_jobs.POST(`delete`, x.Jobs.Delete)
	}
	_schedulers := x.Hertz.Group("schedulers")
	{
		_schedulers.POST(`create`, x.Schedulers.Create)
		_schedulers.POST(`start`, x.Schedulers.Start)
		_schedulers.POST(`stop`, x.Schedulers.Stop)
		_schedulers.POST(`delete`, x.Schedulers.Delete)
	}
	return x.Hertz, nil
}
