package jobs

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/imroc/req/v3"
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
}

type M = map[string]any

func (x *Service) Run(dto CreateDto) (err error) {
	client := req.C().
		SetTimeout(5 * time.Second).
		SetJsonMarshal(sonic.Marshal).
		SetJsonUnmarshal(sonic.Unmarshal)

	if dto.Username != "" && dto.Password != "" {
		client = client.SetCommonBasicAuth(dto.Username, dto.Password)
	}

	r := client.R()
	if dto.Headers != nil {
		r = r.SetHeaders(dto.Headers)
	}
	if dto.Query != nil {
		r = r.SetQueryParams(dto.Query)
	}

	var resp *req.Response
	switch dto.Method {
	case "HEAD":
		if resp, err = r.Head(dto.URL); err != nil {
			return
		}
		break
	case "DELETE":
		if resp, err = r.Delete(dto.URL); err != nil {
			return
		}
		break
	case "POST":
		if dto.Body != "" {
			r = r.SetBodyJsonString(dto.Body)
		}
		if resp, err = r.Post(dto.URL); err != nil {
			return
		}
		break
	case "PATCH":
		if dto.Body != "" {
			r = r.SetBodyJsonString(dto.Body)
		}
		if resp, err = r.Patch(dto.URL); err != nil {
			return
		}
		break
	case "PUT":
		if dto.Body != "" {
			r = r.SetBodyJsonString(dto.Body)
		}
		if resp, err = r.Post(dto.URL); err != nil {
			return
		}
		break
	default:
		if resp, err = r.Get(dto.URL); err != nil {
			return
		}
		break
	}

	println(resp.Status)
	println(resp.StatusCode)
	println(resp.String())

	return
}
