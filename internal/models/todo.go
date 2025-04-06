package models

import (
	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	UserID uint   `gorm:"index;not null"`
	Text   string `gorm:"size:500;not null"`
	Done   bool   `gorm:"default:false"`
}
