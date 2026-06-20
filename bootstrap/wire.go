//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/kainonly/scheduler/api"
	"github.com/kainonly/scheduler/common"

	"github.com/google/wire"
)

func NewAPI(values *common.Values) (*api.API, error) {
	wire.Build(
		wire.Struct(new(api.API), "*"),
		wire.Struct(new(common.Inject), "*"),
		UseScheduler,
		UseHertz,
		api.Provides,
	)
	return &api.API{}, nil
}
