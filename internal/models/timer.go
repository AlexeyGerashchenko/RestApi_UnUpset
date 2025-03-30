package models

import (
	"gorm.io/gorm"
	"time"
)

type Timer struct {
	gorm.Model
	UserID   uint
	Duration time.Duration
}
