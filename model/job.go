package model

import (
	"database/sql/driver"

	"github.com/bytedance/sonic"
)

type Job struct {
	ID          string     `gorm:"primaryKey"`
	SchedulerID string     `gorm:"not null;index"`
	Crontab     string     `gorm:"type:varchar;not null"`
	Schema      *JobSchema `gorm:"type:text;not null"`
}

type JobSchema struct {
	Method   string            `json:"method" vd:"required"`
	URL      string            `json:"url" vd:"required"`
	Headers  map[string]string `json:"headers"`
	Query    map[string]string `json:"query"`
	Body     string            `json:"body"`
	Username string            `json:"username"`
	Password string            `json:"password"`
}

func (x *JobSchema) Scan(value interface{}) error {
	return sonic.Unmarshal(value.([]byte), &x)
}

func (x *JobSchema) Value() (driver.Value, error) {
	if x == nil {
		return sonic.Marshal(JobSchema{})
	}
	return sonic.Marshal(x)
}
