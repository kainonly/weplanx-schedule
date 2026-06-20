package jobs

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/kainonly/go/help"
)

type CreateDto struct {
	Key      string            `json:"key" vd:"required"`
	UUID     uuid.UUID         `json:"uuid" vd:"required"`
	Crontab  string            `json:"crontab" vd:"required"`
	Method   string            `json:"method" vd:"required"`
	URL      string            `json:"url" vd:"required"`
	Headers  map[string]string `json:"headers"`
	Query    map[string]string `json:"query"`
	Body     string            `json:"body"`
	Username string            `json:"username"`
	Password string            `json:"password"`
}

func (x *Controller) Create(ctx context.Context, c *app.RequestContext) {
	var dto CreateDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.JobsX.Create(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Create(ctx context.Context, dto CreateDto) (err error) {
	if !x.Cron.Has(dto.Key) {
		return
	}
	if _, err = x.Cron.Get(dto.Key).NewJob(
		gocron.CronJob(dto.Crontab, true),
		gocron.NewTask(x.Run, dto),
		gocron.WithIdentifier(dto.UUID),
	); err != nil {
		return
	}

	return
}
