package schedulers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/kainonly/go/help"
)

type StartDto struct {
	Key string `json:"key" vd:"required"`
}

func (x *Controller) Start(ctx context.Context, c *app.RequestContext) {
	var dto StartDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.SchedulersX.Start(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Start(ctx context.Context, dto StartDto) (err error) {
	if !x.Cron.Has(dto.Key) {
		return
	}
	s := x.Cron.Get(dto.Key)
	s.Start()
	return
}
