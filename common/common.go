package common

import (
	badger "github.com/dgraph-io/badger/v4"
)

type Inject struct {
	V    *Values
	Db   *badger.DB
	Cron *Cronx
}

type Scheduler struct {
	Key      string          `json:"key" vd:"uuid4"`
	Status   *bool           `json:"status" vd:"required"`
	Name     string          `json:"name" vd:"required"`
	Timezone string          `json:"timezone" vd:"required"`
	Jobs     map[string]*Job `json:"jobs"`
}

type Job struct {
	Identifier string            `json:"identifier" vd:"uuid4"`
	Crontab    string            `json:"crontab" vd:"required"`
	Method     string            `json:"method" vd:"required"`
	URL        string            `json:"url" vd:"required"`
	Headers    map[string]string `json:"headers"`
	Query      map[string]string `json:"query"`
	Body       string            `json:"body"`
	Username   string            `json:"username"`
	Password   string            `json:"password"`
}
