package models

import (
	"time"

	"gorm.io/gorm"
)

// Statistics представляет модель для хранения статистических данных пользователя
type Statistics struct {
	gorm.Model                   // Встраиваем базовую модель GORM для полей ID, CreatedAt, UpdatedAt, DeletedAt
	UserID         uint          `gorm:"index;not null"` // Идентификатор пользователя, к которому относится статистика
	CompletedTasks int           `gorm:"default:0"`      // Количество завершенных задач, по умолчанию 0
	FocusDuration  time.Duration `gorm:"default:0"`      // Общая продолжительность концентрации (работы с таймером)
}
