package storage

import (
	"github.com/bytedance/sonic"
	"github.com/dgraph-io/badger/v4"
	"github.com/kainonly/cronx/common"
)

func (x *Service) SetValue(txn *badger.Txn, key string, data common.Scheduler) (err error) {
	var b []byte
	if b, err = sonic.Marshal(data); err != nil {
		return
	}
	if err = txn.Set([]byte(key), b); err != nil {
		return
	}
	return
}
