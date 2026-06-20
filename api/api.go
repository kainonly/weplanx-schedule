package api

import (
	"context"

	"github.com/kainonly/cronx/api/index"
	"github.com/kainonly/cronx/common"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/google/wire"
)

var Provides = wire.NewSet(
	index.Provides,
)

type API struct {
	*common.Inject

	Hertz  *server.Hertz
	Index  *index.Controller
	IndexX *index.Service
}

func (x *API) Initialize(ctx context.Context) (h *server.Hertz, err error) {

	x.Hertz.GET("", x.Index.Ping)

	return x.Hertz, nil
}
