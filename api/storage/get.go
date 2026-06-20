package storage

import (
	"github.com/bytedance/sonic"
	"github.com/dgraph-io/badger/v4"
	"github.com/kainonly/cronx/common"
)

func (x *Service) Get(txn *badger.Txn, key string) (data common.Scheduler, err error) {
	var it *badger.Item
	if it, err = txn.Get([]byte(key)); err != nil {
		return
	}
	var b []byte
	if b, err = it.ValueCopy(nil); err != nil {
		return
	}
	if len(b) == 0 {
		err = common.ErrStorageNotExists
		return
	}
	if err = sonic.Unmarshal(b, &data); err != nil {
		return
	}
	return
}
