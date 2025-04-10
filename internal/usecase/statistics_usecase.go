package usecase

import (
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
	"time"
)

// StatisticsUC реализует бизнес-логику для работы со статистикой
type StatisticsUC struct {
	statisticsRepo repository.StatisticsRepository // Репозиторий для работы с данными статистики
}

// NewStatisticsUC создает новый экземпляр сервиса статистики
func NewStatisticsUC(statisticsRepo repository.StatisticsRepository) *StatisticsUC {
	return &StatisticsUC{statisticsRepo}
}

// Create добавляет новую запись статистики в базу данных
func (s StatisticsUC) Create(statistics *models.Statistics) error {
	return s.statisticsRepo.Create(statistics)
}

// Update обновляет информацию о статистике в базе данных
func (s StatisticsUC) Update(statistics *models.Statistics) error {
	return s.statisticsRepo.Update(statistics)
}

// Delete удаляет статистику по её идентификатору
func (s StatisticsUC) Delete(id uint) error {
	return s.statisticsRepo.Delete(id)
}

// GetByID возвращает статистику по её идентификатору
func (s StatisticsUC) GetByID(id uint) (*models.Statistics, error) {
	return s.statisticsRepo.GetByID(id)
}

// GetByUserID возвращает статистику конкретного пользователя
func (s StatisticsUC) GetByUserID(userID uint) (*models.Statistics, error) {
	return s.statisticsRepo.GetByUserID(userID)
}

// IncrementCompletedTasks увеличивает счетчик выполненных задач пользователя на 1
func (s StatisticsUC) IncrementCompletedTasks(userID uint) error {
	return s.statisticsRepo.IncrementCompletedTasks(userID)
}

// AddFocusDuration добавляет указанное время к общему времени фокусировки пользователя
func (s StatisticsUC) AddFocusDuration(userID uint, duration time.Duration) error {
	return s.statisticsRepo.AddFocusDuration(userID, duration)
}
