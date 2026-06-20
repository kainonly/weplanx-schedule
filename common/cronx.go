package common

import (
	"sync"

	"github.com/go-co-op/gocron/v2"
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
	return nil, ErrConfigNotExists
}

func (x *Cronx) Remove(key string) (err error) {
	var scheduler gocron.Scheduler
	if scheduler, err = x.Get(key); err != nil {
		return
	}
	if err = scheduler.Shutdown(); err != nil {
		return
	}
	x.m.Delete(key)
	return
}
