package schedulers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/badger/v4"
	"github.com/go-co-op/gocron/v2"
	"github.com/kainonly/go/help"
)

type StopDto struct {
	Key string `json:"key" vd:"uuid4"`
}

func (x *Controller) Stop(ctx context.Context, c *app.RequestContext) {
	var dto StopDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.SchedulersX.Stop(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Stop(ctx context.Context, dto StopDto) error {
	return x.Db.View(func(txn *badger.Txn) (err error) {
		if _, err = x.ConfigsX.Get(txn, dto.Key); err != nil {
			return
		}

		var s gocron.Scheduler
		if s, err = x.Cron.Get(dto.Key); err != nil {
			return
		}

		return s.StopJobs()
	})

}
