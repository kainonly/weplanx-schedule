package schedulers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/badger/v4"
	"github.com/kainonly/go/help"
)

type DeleteDto struct {
	Key string `path:"key" vd:"uuid4"`
}

func (x *Controller) Delete(ctx context.Context, c *app.RequestContext) {
	var dto DeleteDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.SchedulersX.Delete(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Delete(ctx context.Context, dto DeleteDto) error {
	return x.Db.Update(func(txn *badger.Txn) (err error) {
		if _, err = x.StorageX.GetValue(txn, dto.Key); err != nil {
			return
		}

		if _, err = x.Cron.Get(dto.Key); err != nil {
			return
		}

		if err = x.Cron.Remove(dto.Key); err != nil {
			return
		}

		return x.StorageX.Remove(txn, dto.Key)
	})
}
