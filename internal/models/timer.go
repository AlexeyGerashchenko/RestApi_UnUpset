package models

import (
	"time"

	"gorm.io/gorm"
)

// Timer представляет модель таймера для отслеживания времени выполнения задач
type Timer struct {
	gorm.Model               // Встраиваем базовую модель GORM для полей ID, CreatedAt, UpdatedAt, DeletedAt
	UserID     uint          `gorm:"index;not null"` // Идентификатор пользователя, владеющего таймером
	Duration   time.Duration `gorm:"not null"`       // Продолжительность таймера в наносекундах (согласно типу time.Duration)
}
