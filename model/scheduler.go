package model

type Scheduler struct {
	ID       string `gorm:"primaryKey"`
	Status   *bool  `gorm:"not null;default:1"`
	Name     string `gorm:"type:varchar;not null;uniqueIndex"`
	Timezone string `gorm:"type:varchar;not null"`
}
