package configs

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/badger/v4"
	"github.com/kainonly/cronx/common"
)

func (x *Controller) List(ctx context.Context, c *app.RequestContext) {
	results, err := x.ConfigsX.List(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, results)
}

func (x *Service) List(ctx context.Context) (results []common.Scheduler, err error) {
	results = make([]common.Scheduler, 0)
	if err = x.Db.View(func(txn *badger.Txn) (errX error) {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			var b []byte
			if b, errX = item.ValueCopy(nil); errX != nil {
				return
			}

			var data common.Scheduler
			if err = sonic.Unmarshal(b, &data); err != nil {
				return
			}

			results = append(results, data)
		}
		return
	}); err != nil {
		return
	}
	return
}
