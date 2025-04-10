// Package dto содержит структуры для передачи данных между слоями приложения
package dto

// Response представляет общий формат ответа API
type Response struct {
	Status  string      `json:"status"`            // Статус ответа (success/error)
	Message string      `json:"message,omitempty"` // Сообщение (обычно используется для ошибок)
	Data    interface{} `json:"data,omitempty"`    // Данные ответа (опционально)
}

// NewSuccessResponse создает новый успешный ответ с данными
// Принимает произвольные данные и возвращает структуру Response с заполненными полями
func NewSuccessResponse(data interface{}) Response {
	return Response{
		Status: "success",
		Data:   data,
	}
}

// NewErrorResponse создает новый ответ с ошибкой
// Принимает сообщение об ошибке и возвращает структуру Response с заполненными полями
func NewErrorResponse(message string) Response {
	return Response{
		Status:  "error",
		Message: message,
	}
}
