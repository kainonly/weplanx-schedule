package schedulers

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/kainonly/cronx/model"
	"github.com/kainonly/go/help"
	"gorm.io/gorm"
)

type CreateDto struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"name" vd:"required"`
	Timezone string    `json:"timezone" vd:"required"`
}

func (x *Controller) Create(ctx context.Context, c *app.RequestContext) {
	var dto CreateDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	dto.ID = uuid.New()
	if err := x.SchedulersX.Create(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, utils.H{
		"uuid": dto.ID,
	})
}

func (x *Service) Create(ctx context.Context, dto CreateDto) (err error) {
	var exists int64
	if err = x.Db.Model(model.Scheduler{}).WithContext(ctx).
		Where("name = ?", dto.Name).
		Count(&exists).Error; err != nil {
		return
	}

	if exists != 0 {
		return help.E(0, `The [name] already exists.`)
	}

	return x.Db.Transaction(func(tx *gorm.DB) (errX error) {
		data := model.Scheduler{
			ID:       dto.ID.String(),
			Name:     dto.Name,
			Timezone: dto.Timezone,
		}
		if errX = tx.WithContext(ctx).
			Create(&data).Error; errX != nil {
			return
		}

		var tz *time.Location
		if tz, err = time.LoadLocation(data.Timezone); err != nil {
			return
		}
		var s gocron.Scheduler
		if s, err = gocron.NewScheduler(
			gocron.WithLocation(tz),
		); err != nil {
			return
		}
		x.Cron.Store(dto.ID.String(), s)
		return
	})
}
