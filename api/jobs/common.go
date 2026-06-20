package jobs

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
	"github.com/kainonly/cronx/api/schedulers"
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

	SchedulersX *schedulers.Service
}

type M = map[string]any

func (x *Service) Run(dto CreateDto) (err error) {
	client := req.C().
		SetTimeout(5 * time.Second).
		SetJsonMarshal(sonic.Marshal).
		SetJsonUnmarshal(sonic.Unmarshal)

	if dto.Schema.Username != "" && dto.Schema.Password != "" {
		client = client.SetCommonBasicAuth(
			dto.Schema.Username,
			dto.Schema.Password,
		)
	}

	r := client.R()
	if dto.Schema.Headers != nil {
		r = r.SetHeaders(dto.Schema.Headers)
	}
	if dto.Schema.Query != nil {
		r = r.SetQueryParams(dto.Schema.Query)
	}

	var resp *req.Response
	switch dto.Schema.Method {
	case "HEAD":
		if resp, err = r.Head(dto.Schema.URL); err != nil {
			return
		}
		break
	case "DELETE":
		if resp, err = r.Delete(dto.Schema.URL); err != nil {
			return
		}
		break
	case "POST":
		if dto.Schema.Body != "" {
			r = r.SetBodyJsonString(dto.Schema.Body)
		}
		if resp, err = r.Post(dto.Schema.URL); err != nil {
			return
		}
		break
	case "PATCH":
		if dto.Schema.Body != "" {
			r = r.SetBodyJsonString(dto.Schema.Body)
		}
		if resp, err = r.Patch(dto.Schema.URL); err != nil {
			return
		}
		break
	case "PUT":
		if dto.Schema.Body != "" {
			r = r.SetBodyJsonString(dto.Schema.Body)
		}
		if resp, err = r.Post(dto.Schema.URL); err != nil {
			return
		}
		break
	default:
		if resp, err = r.Get(dto.Schema.URL); err != nil {
			return
		}
		break
	}

	println(resp.Status)
	println(resp.StatusCode)
	println(resp.String())

	return
}
