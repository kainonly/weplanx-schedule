package jobs

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/kainonly/go/help"
)

type DeleteDto struct {
	Key  string    `json:"key" vd:"required"`
	UUID uuid.UUID `json:"uuid" vd:"required"`
}

func (x *Controller) Delete(ctx context.Context, c *app.RequestContext) {
	var dto DeleteDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.JobsX.Delete(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Delete(ctx context.Context, dto DeleteDto) (err error) {
	if !x.Cron.Has(dto.Key) {
		return
	}
	return x.Cron.Get(dto.Key).RemoveJob(dto.UUID)
}
