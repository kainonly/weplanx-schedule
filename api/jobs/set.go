package jobs

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/badger/v4"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/kainonly/cronx/common"
	"github.com/kainonly/go/help"
)

type SetDto struct {
	SchedulerKey string `json:"schedule_key" vd:"required,uuid4"`
	*common.Job
}

func (x *Controller) Set(ctx context.Context, c *app.RequestContext) {
	var dto SetDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.JobsX.Set(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) SetRunner(key string, job *common.Job) (err error) {
	var identifier uuid.UUID
	if identifier, err = uuid.FromBytes([]byte(job.Identifier)); err != nil {
		return
	}

	var scheduler gocron.Scheduler
	if scheduler, err = x.Cron.Get(key); err != nil {
		return
	}

	if _, err = scheduler.NewJob(
		gocron.CronJob(job.Crontab, true),
		gocron.NewTask(x.Run, job),
		gocron.WithIdentifier(identifier),
	); err != nil {
		return
	}
	return
}

func (x *Service) Set(ctx context.Context, dto SetDto) error {
	return x.Db.Update(func(txn *badger.Txn) (err error) {
		var data common.Scheduler
		if data, err = x.StorageX.Get(txn, dto.SchedulerKey); err != nil {
			return
		}

		if err = x.SetRunner(dto.SchedulerKey, dto.Job); err != nil {
			return
		}

		data.Jobs[dto.Identifier] = dto.Job
		return x.StorageX.Set(txn, dto.SchedulerKey, data)
	})
}
