package models

import (
	"gorm.io/gorm"
	"time"
)

// Храню статистику по дням - каждый день отдельно

type Statistics struct {
	gorm.Model
	UserID         uint          `gorm:"index;not null"`
	Date           time.Time     `gorm:"type:date;index;not null"`
	CompletedTasks int           `gorm:"default:0"`
	FocusDuration  time.Duration `gorm:"default:0"`
}
