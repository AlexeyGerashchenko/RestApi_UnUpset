package dto

// CreateToDoRequest представляет данные для создания новой задачи
type CreateToDoRequest struct {
	Text string `json:"text" binding:"required,min=1,max=500"` // Текст задачи (обязательное поле, от 1 до 500 символов)
}

// ToDorResponse представляет данные задачи для ответа API
type ToDorResponse struct {
	ID     uint   `json:"id"`      // Идентификатор задачи
	UserID uint   `json:"user_id"` // Идентификатор пользователя, создавшего задачу
	Text   string `json:"text"`    // Текст задачи
	Done   bool   `json:"done"`    // Статус выполнения задачи (true - выполнена, false - не выполнена)
}
