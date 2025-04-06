package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName   string       `gorm:"size:100;not null"`
	Email      string       `gorm:"size:100;uniqueIndex;not null"`
	Password   string       `gorm:"size:255;not null"`
	ToDos      []ToDo       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Timers     []Timer      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Statistics []Statistics `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
