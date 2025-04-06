package models

import (
	"gorm.io/gorm"
	"time"
)

type Timer struct {
	gorm.Model
	UserID   uint          `gorm:"index;not null"`
	Duration time.Duration `gorm:"not null"`
}
