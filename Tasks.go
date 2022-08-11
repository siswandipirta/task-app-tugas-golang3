package Tasks

import (
	"time"

	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	Id       int `gorm:"primaryKey;autoIncrement:true"`
	Task     string
	Assign   string
	Deadline *time.Time
	Status   string
}
