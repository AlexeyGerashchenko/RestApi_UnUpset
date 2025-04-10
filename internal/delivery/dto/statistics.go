package dto

// StatisticsResponse представляет данные статистики для ответа API
type StatisticsResponse struct {
	CompletedTasks int    `json:"completed_tasks"` // Количество выполненных задач
	FocusDuration  string `json:"focus_duration"`  // Общая продолжительность фокусировки в читаемом формате
}
