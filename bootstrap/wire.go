//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/kainonly/cronx/api"
	"github.com/kainonly/cronx/common"

	"github.com/google/wire"
)

func NewAPI(values *common.Values) (*api.API, error) {
	wire.Build(
		wire.Struct(new(api.API), "*"),
		wire.Struct(new(common.Inject), "*"),
		UseCronx,
		UseHertz,
		api.Provides,
	)
	return &api.API{}, nil
}
