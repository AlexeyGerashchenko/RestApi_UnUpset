// Package models содержит структуры данных, описывающие модели приложения
package models

import "gorm.io/gorm"

// User представляет модель пользователя в системе
type User struct {
	gorm.Model              // Встраиваем базовую модель GORM для полей ID, CreatedAt, UpdatedAt, DeletedAt
	UserName   string       `gorm:"size:50;not null"`                              // Имя пользователя, ограничено 50 символами
	Email      string       `gorm:"size:100;uniqueIndex;not null"`                 // Email, уникальный индекс, до 100 символов
	Password   string       `gorm:"size:255;not null"`                             // Хеш пароля, до 255 символов
	ToDos      []ToDo       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Связь один-ко-многим с задачами, каскадное удаление
	Timers     []Timer      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Связь один-ко-многим с таймерами, каскадное удаление
	Statistics []Statistics `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Связь один-ко-многим со статистикой, каскадное удаление
}
