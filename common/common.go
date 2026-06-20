package common

import (
	"sync"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

type Inject struct {
	V    *Values
	Db   *gorm.DB
	Cron *Cronx
}

type Cronx struct {
	m sync.Map
}

func (x *Cronx) Store(id string, value gocron.Scheduler) {
	x.m.Store(id, value)
}

func (x *Cronx) Has(id string) bool {
	if _, ok := x.m.Load(id); ok {
		return ok
	}
	return false
}

func (x *Cronx) Get(id string) (value gocron.Scheduler) {
	if v, ok := x.m.Load(id); ok {
		return v.(gocron.Scheduler)
	}
	return
}

func (x *Cronx) Remove(id string) (err error) {
	s := x.Get(id)
	if err = s.Shutdown(); err != nil {
		return
	}
	x.m.Delete(id)
	return
}
