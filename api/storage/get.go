package storage

import (
	"github.com/bytedance/sonic"
	"github.com/dgraph-io/badger/v4"
	"github.com/kainonly/cronx/common"
)

func (x *Service) Get(txn *badger.Txn, key string) (data common.Scheduler, err error) {
	var item *badger.Item
	if item, err = txn.Get([]byte(key)); err != nil {
		return
	}
	var b []byte
	if b, err = item.ValueCopy(nil); err != nil {
		return
	}
	if len(b) == 0 {
		err = common.ErrConfigNotExists
		return
	}
	if err = sonic.Unmarshal(b, &data); err != nil {
		return
	}
	return
}
