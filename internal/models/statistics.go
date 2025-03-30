package models

import (
	"gorm.io/gorm"
	"time"
)

type Statistics struct {
	gorm.Model
	UserID         uint
	Date           time.Time
	CompletedTasks int
	FocusDuration  time.Duration
}
