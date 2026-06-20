package common

import (
	"github.com/bytedance/sonic"
	badger "github.com/dgraph-io/badger/v4"
	"github.com/kainonly/go/help"
	"resty.dev/v3"
)

type Inject struct {
	V    *Values
	Db   *badger.DB
	Logs *Victorialogs
	Cron *Cronx
}

var (
	ErrNotExists       = help.E(0, `The key does not exist in the process`)
	ErrConfigNotExists = help.E(0, `The key does not exist in the config.`)
)

type Scheduler struct {
	Key      string          `json:"key" vd:"required,uuid4"`
	Status   *bool           `json:"status" vd:"required"`
	Name     string          `json:"name" vd:"required"`
	Timezone string          `json:"timezone" vd:"required,timezone"`
	Jobs     map[string]*Job `json:"jobs" vd:"required,dive,keys,uuid4,endkeys,required"`
}

type Job struct {
	SchedulerKey string            `json:"scheduler_key" vd:"required,uuid4"`
	Id           string            `json:"id" vd:"required,uuid4"`
	Crontab      string            `json:"crontab" vd:"required,cron"`
	Method       string            `json:"method" vd:"oneof=GET HEAD DELETE POST PATCH PUT"`
	URL          string            `json:"url" vd:"url"`
	Headers      map[string]string `json:"headers"`
	Query        map[string]string `json:"query"`
	Body         string            `json:"body"`
	Username     string            `json:"username"`
	Password     string            `json:"password"`
}

type Victorialogs struct {
	Client *resty.Client
}

type PushDto struct {
	Ts           string  `json:"ts"` // ISO8601 or RFC3339
	SchedulerKey string  `json:"scheduler_key"`
	JobId        string  `json:"job_id"`
	Log          PushLog `json:"log"`
}

type PushLog struct {
	Duration string `json:"duration"`
	Status   string `json:"status"`
	Body     string `json:"body"`
}

func (x *Victorialogs) Push(dto PushDto) (err error) {
	var body string
	if body, err = sonic.MarshalString(dto); err != nil {
		return
	}
	var resp *resty.Response
	if resp, err = x.Client.R().
		SetQueryParam(`_time_field`, `ts`).
		SetQueryParam(`_stream_fields`, `scheduler_key,job_id`).
		SetQueryParam(`_msg_field`, `log.body`).
		SetBody(body).
		Post(`/insert/jsonline`); err != nil {
		return
	}
	if resp.IsError() {
		return
	}
	return
}
