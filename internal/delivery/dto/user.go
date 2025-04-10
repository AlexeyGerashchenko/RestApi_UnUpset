package dto

import "time"

// LoginRequest представляет данные для авторизации пользователя
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"` // Email пользователя (обязательное поле)
	Password string `json:"password" binding:"required"`    // Пароль пользователя (обязательное поле)
}

// RegisterUserRequest представляет данные для регистрации нового пользователя
type RegisterUserRequest struct {
	UserName string `json:"username" binding:"required,min=3,max=100"` // Имя пользователя (обязательное поле, от 3 до 100 символов)
	Email    string `json:"email" binding:"required,email"`            // Email пользователя (обязательное поле)
	Password string `json:"password" binding:"required,min=5"`         // Пароль пользователя (обязательное поле, минимум 5 символов)
}

// UserResponse представляет данные пользователя для ответа API
type UserResponse struct {
	ID        uint      `json:"id"`         // Идентификатор пользователя
	UserName  string    `json:"username"`   // Имя пользователя
	Email     string    `json:"email"`      // Email пользователя
	CreatedAt time.Time `json:"created_at"` // Дата создания учетной записи
}

// ChangePasswordRequest представляет данные для изменения пароля пользователя
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`       // Текущий пароль пользователя (обязательное поле)
	NewPassword string `json:"new_password" binding:"required,min=5"` // Новый пароль пользователя (обязательное поле, минимум 5 символов)
}

// ChangeUsernameRequest представляет данные для изменения имени пользователя
type ChangeUsernameRequest struct {
	NewUsername string `json:"new_username" binding:"required,min=3,max=50"` // Новое имя пользователя (обязательное поле, от 3 до 50 символов)
}
