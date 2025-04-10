package usecase

import (
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
)

// TimerUC реализует бизнес-логику для работы с таймерами
type TimerUC struct {
	timerRepo    repository.TimerRepository // Репозиторий для работы с данными таймеров
	statisticsUC StatisticsUC               // Сервис для работы со статистикой
}

// NewTimerUC создает новый экземпляр сервиса таймеров
func NewTimerUC(timerRepo repository.TimerRepository, statisticsUC StatisticsUC) *TimerUC {
	return &TimerUC{
		timerRepo:    timerRepo,
		statisticsUC: statisticsUC,
	}
}

// Create добавляет новый таймер и обновляет статистику времени фокусировки
func (t TimerUC) Create(timer *models.Timer) error {
	// Создаем запись таймера в базе данных
	if err := t.timerRepo.Create(timer); err != nil {
		return err
	}

	// Добавляем время фокусировки в статистику пользователя
	return t.statisticsUC.AddFocusDuration(timer.UserID, timer.Duration)
}

// Delete удаляет таймер по его идентификатору
func (t TimerUC) Delete(id uint) error {
	return t.timerRepo.Delete(id)
}

// GetByID возвращает таймер по его идентификатору
func (t TimerUC) GetByID(id uint) (*models.Timer, error) {
	return t.timerRepo.GetByID(id)
}

// GetByUserID возвращает все таймеры конкретного пользователя
func (t TimerUC) GetByUserID(userID uint) ([]*models.Timer, error) {
	return t.timerRepo.GetByUserID(userID)
}
