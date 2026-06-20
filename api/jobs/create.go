package jobs

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/kainonly/cronx/model"
	"github.com/kainonly/go/help"
	"gorm.io/gorm"
)

type CreateDto struct {
	SchedulerID string           `json:"schedule_id" vd:"required,uuid4"`
	Crontab     string           `json:"crontab" vd:"required"`
	Schema      *model.JobSchema `json:"schema" vd:"required,dive"`
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
	if err = x.SchedulersX.CheckSchedulerExists(ctx, dto.SchedulerID); err != nil {
		return
	}
	if !x.Cron.Has(dto.SchedulerID) {
		return
	}

	return x.Db.Transaction(func(tx *gorm.DB) (errX error) {
		jobID := uuid.New()
		data := model.Job{
			ID:          jobID.String(),
			SchedulerID: dto.SchedulerID,
			Crontab:     dto.Crontab,
			Schema:      dto.Schema,
		}

		if errX = tx.WithContext(ctx).
			Create(&data).Error; errX != nil {
			return
		}

		if _, err = x.Cron.Get(dto.SchedulerID).NewJob(
			gocron.CronJob(dto.Crontab, true),
			gocron.NewTask(x.Run, dto),
			gocron.WithIdentifier(jobID),
		); err != nil {
			return
		}

		return
	})
}
