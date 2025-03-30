package models

import (
	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	UserID uint
	Text   string
	Done   bool
}
