package schedulers

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/kainonly/cronx/model"
	"github.com/kainonly/go/help"
	"gorm.io/gorm"
)

type DeleteDto struct {
	ID string `json:"id" vd:"required,uuid4"`
}

func (x *Controller) Delete(ctx context.Context, c *app.RequestContext) {
	var dto DeleteDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.SchedulersX.Delete(ctx, dto); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, help.Ok())
}

func (x *Service) Delete(ctx context.Context, dto DeleteDto) (err error) {
	if err = x.CheckSchedulerExists(ctx, dto.ID); err != nil {
		return
	}

	return x.Db.Transaction(func(tx *gorm.DB) (errX error) {
		if errX = tx.Model(model.Scheduler{}).WithContext(ctx).
			Delete(dto.ID).Error; errX != nil {
			return
		}

		if !x.Cron.Has(dto.ID) {
			return
		}
		return x.Cron.Remove(dto.ID)
	})

}
