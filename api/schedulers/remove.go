package schedulers

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/badger/v4"
	"github.com/kainonly/cronx/common"
	"github.com/kainonly/go/help"
)

type RemoveDto struct {
	Key string `path:"key" vd:"uuid4"`
}

func (x *Controller) Remove(ctx context.Context, c *app.RequestContext) {
	var dto RemoveDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.SchedulersX.Remove(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Remove(ctx context.Context, dto RemoveDto) error {
	return x.Db.Update(func(txn *badger.Txn) (err error) {
		if err = x.Cron.Remove(dto.Key); err != nil && !errors.Is(err, common.ErrConfigNotExists) {
			return
		}
		return x.StorageX.Remove(txn, dto.Key)
	})
}
