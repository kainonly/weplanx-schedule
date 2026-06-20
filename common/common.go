package common

import (
	badger "github.com/dgraph-io/badger/v4"
	"github.com/kainonly/go/help"
)

type Inject struct {
	V    *Values
	Db   *badger.DB
	Cron *Cronx
}

var (
	ErrNotExists       = help.E(0, `The key does not exist in the process`)
	ErrConfigNotExists = help.E(0, `The key does not exist in the config.`)
)

type Scheduler struct {
	Key      string          `json:"key" vd:"uuid4"`
	Status   *bool           `json:"status" vd:"required"`
	Name     string          `json:"name" vd:"required"`
	Timezone string          `json:"timezone" vd:"timezone"`
	Jobs     map[string]*Job `json:"jobs" vd:"required,dive,keys,uuid4,endkeys,required"`
}

type Job struct {
	Id       string            `json:"id" vd:"uuid4"`
	Crontab  string            `json:"crontab" vd:"cron"`
	Method   string            `json:"method" vd:"oneof=GET HEAD DELETE POST PATCH PUT"`
	URL      string            `json:"url" vd:"url"`
	Headers  map[string]string `json:"headers"`
	Query    map[string]string `json:"query"`
	Body     string            `json:"body"`
	Username string            `json:"username"`
	Password string            `json:"password"`
}

type Victorialogs struct {
}
