package jobs

import (
	"time"

	"github.com/kainonly/cronx/api/configs"
	"github.com/kainonly/cronx/common"
	"resty.dev/v3"

	"github.com/google/wire"
)

var Provides = wire.NewSet(
	wire.Struct(new(Controller), "*"),
	wire.Struct(new(Service), "*"),
)

type Controller struct {
	V *common.Values

	JobsX *Service
}

type Service struct {
	*common.Inject

	ConfigsX *configs.Service
}

type M = map[string]any

func (x *Service) Run(cfg common.Job) (err error) {
	client := resty.New().
		SetTimeout(5 * time.Second)

	if cfg.Username != "" && cfg.Password != "" {
		client.SetBasicAuth(cfg.Username, cfg.Password)
	}

	r := client.R()
	if cfg.Headers != nil {
		r = r.SetHeaders(cfg.Headers)
	}
	if cfg.Query != nil {
		r = r.SetQueryParams(cfg.Query)
	}

	var resp *resty.Response
	switch cfg.Method {
	case "HEAD":
		if resp, err = r.Head(cfg.URL); err != nil {
			return
		}
	case "DELETE":
		if resp, err = r.Delete(cfg.URL); err != nil {
			return
		}
	case "POST":
		if cfg.Body != "" {
			r = r.SetHeader("Content-Type", "application/json").SetBody(cfg.Body)
		}
		if resp, err = r.Post(cfg.URL); err != nil {
			return
		}
	case "PATCH":
		if cfg.Body != "" {
			r = r.SetHeader("Content-Type", "application/json").SetBody(cfg.Body)
		}
		if resp, err = r.Patch(cfg.URL); err != nil {
			return
		}
	case "PUT":
		if cfg.Body != "" {
			r = r.SetHeader("Content-Type", "application/json").SetBody(cfg.Body)
		}
		if resp, err = r.Put(cfg.URL); err != nil {
			return
		}
	default:
		if resp, err = r.Get(cfg.URL); err != nil {
			return
		}
	}

	println(resp.Status())
	println(resp.StatusCode())
	println(resp.String())

	return
}

func (x *Service) Push(cfg common.Job) (err error) {
	//resp, err := client.R().
	//	AddQueryParam(`_stream_fields`, `schedule_id,job_id`).
	//	AddQueryParam(`_time_field`, `date`).
	//	AddQueryParam(`_msg_field`, `log.message`).
	//	SetBodyString(`
	//		{ "log": { "level": "info", "message": "hello world" }, "date": "0", "schedule_id": "schedule1", "job_id": "a" }
	//		{ "log": { "level": "error", "message": "oh no!" }, "date": "0", "schedule_id": "schedule1", "job_id": "b" }
	//		{ "log": { "level": "info", "message": "hello world" }, "date": "0", "schedule_id": "schedule2", "job_id": "a" }`).
	//	Post(`/insert/jsonline`)

	return
}
