package jobs

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/badger/v4"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/kainonly/go/help"
)

type RemoveDto struct {
	SchedulerKey string `json:"schedule_key" vd:"uuid4"`
	Identifier   string `json:"identifier" vd:"uuid4"`
}

func (x *Controller) Remove(ctx context.Context, c *app.RequestContext) {
	var dto RemoveDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.JobsX.Delete(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Delete(ctx context.Context, dto RemoveDto) error {
	return x.Db.Update(func(txn *badger.Txn) (err error) {
		if _, err = x.StorageX.GetValue(txn, dto.SchedulerKey); err != nil {
			return
		}

		var identifier uuid.UUID
		if identifier, err = uuid.FromBytes([]byte(dto.Identifier)); err != nil {
			return
		}

		var scheduler gocron.Scheduler
		if scheduler, err = x.Cron.Get(dto.SchedulerKey); err != nil {
			return
		}

		if err = scheduler.RemoveJob(identifier); err != nil {
			return
		}

		// TODO: 合并配置再更新本地存储...
		return
	})
}
