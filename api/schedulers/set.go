package schedulers

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/badger/v4"
	"github.com/go-co-op/gocron/v2"
	"github.com/kainonly/cronx/common"
	"github.com/kainonly/go/help"
)

type SetDto struct {
	*common.Scheduler
}

func (x *Controller) Set(ctx context.Context, c *app.RequestContext) {
	var dto SetDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.SchedulersX.Set(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Set(ctx context.Context, dto SetDto) error {
	return x.Db.Update(func(txn *badger.Txn) (err error) {
		var scheduler gocron.Scheduler
		if scheduler, err = x.Cron.Get(dto.Key); err != nil && !errors.Is(err, common.ErrNotExists) {
			return
		}
		if scheduler != nil {
			if err = scheduler.Shutdown(); err != nil {
				return
			}
		}

		var tz *time.Location
		if tz, err = time.LoadLocation(dto.Timezone); err != nil {
			return
		}

		if scheduler, err = gocron.NewScheduler(gocron.WithLocation(tz)); err != nil {
			return
		}

		x.Cron.Store(dto.Key, scheduler)
		for _, job := range dto.Jobs {
			if err = x.JobsX.SetRunner(dto.Key, *job); err != nil {
				return
			}
		}

		if *dto.Status {
			scheduler.Start()
		}

		return x.ConfigsX.Set(txn, dto.Key, *dto.Scheduler)
	})
}
