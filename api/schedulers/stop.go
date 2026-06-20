package schedulers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/kainonly/go/help"
)

type StopDto struct {
	Key string `json:"key" vd:"required"`
}

func (x *Controller) Stop(ctx context.Context, c *app.RequestContext) {
	var dto StopDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.SchedulersX.Stop(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Stop(ctx context.Context, dto StopDto) (err error) {
	if !x.Cron.Has(dto.Key) {
		return
	}
	s := x.Cron.Get(dto.Key)
	return s.StopJobs()
}
