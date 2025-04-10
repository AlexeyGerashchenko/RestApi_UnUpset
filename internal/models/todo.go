package models

import (
	"gorm.io/gorm"
)

// ToDo представляет модель задачи пользователя в системе
type ToDo struct {
	gorm.Model        // Встраиваем базовую модель GORM для полей ID, CreatedAt, UpdatedAt, DeletedAt
	UserID     uint   `gorm:"index;not null"`    // Идентификатор пользователя, создавшего задачу, индексировано для быстрого поиска
	Text       string `gorm:"size:500;not null"` // Текст задачи, ограничен 500 символами
	Done       bool   `gorm:"default:false"`     // Статус выполнения задачи, по умолчанию false (не выполнена)
}
