package common

import (
	"sync"

	"github.com/go-co-op/gocron/v2"
)

type Inject struct {
	V    *Values
	Cron *Cronx
}

type Cronx struct {
	m sync.Map
}

func (x *Cronx) Store(key string, value gocron.Scheduler) {
	x.m.Store(key, value)
}

func (x *Cronx) Has(key string) bool {
	if _, ok := x.m.Load(key); ok {
		return ok
	}
	return false
}

func (x *Cronx) Get(key string) (value gocron.Scheduler) {
	if v, ok := x.m.Load(key); ok {
		return v.(gocron.Scheduler)
	}
	return
}

func (x *Cronx) Remove(key string) (err error) {
	s := x.Get(key)
	if err = s.Shutdown(); err != nil {
		return
	}
	x.m.Delete(key)
	return
}
