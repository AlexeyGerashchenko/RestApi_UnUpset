package models

import (
	"gorm.io/gorm"
	"time"
)

// Храню статистику по дням - каждый день отдельно

type Statistics struct {
	gorm.Model
	UserID         uint
	Date           time.Time
	CompletedTasks int
	FocusDuration  time.Duration
}
