package schedulers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/kainonly/go/help"
)

type StopDto struct {
	ID string `json:"id" vd:"required,uuid4"`
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
	if err = x.CheckSchedulerExists(ctx, dto.ID); err != nil {
		return
	}
	if !x.Cron.Has(dto.ID) {
		return
	}
	s := x.Cron.Get(dto.ID)
	return s.StopJobs()
}
