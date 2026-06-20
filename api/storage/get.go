package storage

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/badger/v4"
	"github.com/kainonly/cronx/common"
	"github.com/kainonly/go/help"
)

type GetDto struct {
	Key string `path:"key" vd:"uuid4"`
}

func (x *Controller) Get(ctx context.Context, c *app.RequestContext) {

}

func (x *Service) Get(ctx context.Context, dto GetDto) (err error) {
	return
}

func (x *Service) GetValue(txn *badger.Txn, key string) (data common.Scheduler, err error) {
	var it *badger.Item
	if it, err = txn.Get([]byte(key)); err != nil {
		return
	}
	var b []byte
	if b, err = it.ValueCopy(nil); err != nil {
		return
	}
	if len(b) == 0 {
		err = help.E(0, fmt.Sprintf(`The key[%s] does not exist in the storage.`, key))
		return
	}
	if err = sonic.Unmarshal(b, &data); err != nil {
		return
	}
	return
}
