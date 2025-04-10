package dto

import (
	"time"
)

// TimerResponse представляет данные таймера для ответа API
type TimerResponse struct {
	ID       uint          `json:"id"`       // Идентификатор таймера
	UserID   uint          `json:"user_id"`  // Идентификатор пользователя, создавшего таймер
	Duration time.Duration `json:"duration"` // Продолжительность таймера в наносекундах
}

// CreateTimerRequest представляет данные для создания нового таймера
type CreateTimerRequest struct {
	Duration time.Duration `json:"duration" binding:"required,min=1"` // Продолжительность таймера (обязательное поле, минимум 1 наносекунда)
}
