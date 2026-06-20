package common

import (
	"fmt"
	"sync"

	"github.com/go-co-op/gocron/v2"
	"github.com/kainonly/go/help"
)

type Cronx struct {
	m sync.Map
}

func (x *Cronx) Store(key string, value gocron.Scheduler) {
	x.m.Store(key, value)
}

func (x *Cronx) Get(key string) (gocron.Scheduler, error) {
	if v, ok := x.m.Load(key); ok {
		return v.(gocron.Scheduler), nil
	}
	return nil, help.E(0, fmt.Sprintf(`The key[%s] does not exist in the schedule config`, key))
}

func (x *Cronx) Remove(key string) (err error) {
	var s gocron.Scheduler
	if s, err = x.Get(key); err != nil {
		return
	}
	if err = s.Shutdown(); err != nil {
		return
	}
	x.m.Delete(key)
	return
}
