package jobs

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
	"github.com/kainonly/cronx/api/storage"
	"github.com/kainonly/cronx/common"

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

	StorageX *storage.Service
}

type M = map[string]any

func (x *Service) Run(cfg common.Job) (err error) {
	client := req.C().
		SetTimeout(5 * time.Second).
		SetJsonMarshal(sonic.Marshal).
		SetJsonUnmarshal(sonic.Unmarshal)

	if cfg.Username != "" && cfg.Password != "" {
		client = client.SetCommonBasicAuth(cfg.Username, cfg.Password)
	}

	r := client.R()
	if cfg.Headers != nil {
		r = r.SetHeaders(cfg.Headers)
	}
	if cfg.Query != nil {
		r = r.SetQueryParams(cfg.Query)
	}

	var resp *req.Response
	switch cfg.Method {
	case "HEAD":
		if resp, err = r.Head(cfg.URL); err != nil {
			return
		}
		break
	case "DELETE":
		if resp, err = r.Delete(cfg.URL); err != nil {
			return
		}
		break
	case "POST":
		if cfg.Body != "" {
			r = r.SetBodyJsonString(cfg.Body)
		}
		if resp, err = r.Post(cfg.URL); err != nil {
			return
		}
		break
	case "PATCH":
		if cfg.Body != "" {
			r = r.SetBodyJsonString(cfg.Body)
		}
		if resp, err = r.Patch(cfg.URL); err != nil {
			return
		}
		break
	case "PUT":
		if cfg.Body != "" {
			r = r.SetBodyJsonString(cfg.Body)
		}
		if resp, err = r.Post(cfg.URL); err != nil {
			return
		}
		break
	default:
		if resp, err = r.Get(cfg.URL); err != nil {
			return
		}
		break
	}

	println(resp.Status)
	println(resp.StatusCode)
	println(resp.String())

	// TODO: 这里接入 Victoria

	return
}
