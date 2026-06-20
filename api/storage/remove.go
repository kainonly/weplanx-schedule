package storage

import "github.com/dgraph-io/badger/v4"

func (x *Service) Remove(txn *badger.Txn, key string) error {
	return txn.Delete([]byte(key))
}
