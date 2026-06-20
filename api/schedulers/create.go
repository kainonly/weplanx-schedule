package schedulers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/badger/v4"
	"github.com/go-co-op/gocron/v2"
	"github.com/kainonly/cronx/common"
	"github.com/kainonly/go/help"
)

type CreateDto struct {
	*common.Scheduler
}

func (x *Controller) Create(ctx context.Context, c *app.RequestContext) {
	var dto CreateDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.SchedulersX.Create(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Create(ctx context.Context, dto CreateDto) error {
	return x.Db.Update(func(txn *badger.Txn) (err error) {
		if _, err = x.StorageX.GetValue(txn, dto.Key); err != nil {
			return
		}

		var tz *time.Location
		if tz, err = time.LoadLocation(dto.Timezone); err != nil {
			return
		}

		var s gocron.Scheduler
		if s, err = gocron.NewScheduler(
			gocron.WithLocation(tz),
		); err != nil {
			return
		}

		x.Cron.Store(dto.Key, s)
		return x.StorageX.SetValue(txn, dto.Key, *dto.Scheduler)
	})
}
