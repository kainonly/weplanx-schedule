package api

import (
	"context"

	"github.com/kainonly/cronx/api/index"
	"github.com/kainonly/cronx/api/schedulers"
	"github.com/kainonly/cronx/common"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/google/wire"
)

var Provides = wire.NewSet(
	index.Provides,
	schedulers.Provides,
)

type API struct {
	*common.Inject

	Hertz      *server.Hertz
	Index      *index.Controller
	IndexX     *index.Service
	Schedulers *schedulers.Controller
}

func (x *API) Initialize(ctx context.Context) (h *server.Hertz, err error) {
	x.Hertz.GET("", x.Index.Ping)

	_schedulers := x.Hertz.Group("schedulers")
	{
		_schedulers.POST(`create`, x.Schedulers.Create)
		_schedulers.POST(`start`, x.Schedulers.Start)
		_schedulers.POST(`stop`, x.Schedulers.Stop)
		_schedulers.POST(`delete`, x.Schedulers.Delete)
	}
	return x.Hertz, nil
}
