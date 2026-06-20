package schedulers

import (
	"context"

	"github.com/kainonly/cronx/common"
	"github.com/kainonly/cronx/model"
	"github.com/kainonly/go/help"

	"github.com/google/wire"
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
}

type M = map[string]any

func (x *Service) CheckSchedulerExists(ctx context.Context, id string) (err error) {
	var exists int64
	if err = x.Db.Model(model.Scheduler{}).WithContext(ctx).
		Where("id = ?", id).
		Count(&exists).Error; err != nil {
		return
	}

	if exists == 0 {
		return help.E(0, `The scheduler do not exist.`)
	}
	return
}
