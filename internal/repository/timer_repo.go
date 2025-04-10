package repository

import (
	"RestApi_UnUpset/internal/models"

	"gorm.io/gorm"
)

// TimerRepo представляет репозиторий для работы с моделью таймера
type TimerRepo struct {
	db *gorm.DB // Подключение к базе данных
}

// NewTimerRepo создает новый экземпляр репозитория таймеров
func NewTimerRepo(db *gorm.DB) *TimerRepo {
	return &TimerRepo{db: db}
}

// Create добавляет новый таймер в базу данных
func (r *TimerRepo) Create(t *models.Timer) error {
	return r.db.Create(t).Error // Используем GORM для создания записи
}

// GetByID возвращает таймер по его идентификатору
func (r *TimerRepo) GetByID(id uint) (*models.Timer, error) {
	var t models.Timer
	err := r.db.First(&t, id).Error // Ищем первую запись с указанным ID
	return &t, err
}

// GetByUserID возвращает все таймеры конкретного пользователя
func (r *TimerRepo) GetByUserID(userID uint) ([]*models.Timer, error) {
	var timers []*models.Timer
	// Фильтруем по ID пользователя
	err := r.db.Where("user_id = ?", userID).Find(&timers).Error
	return timers, err
}

// Delete удаляет таймер из базы данных по ID
// Используется soft delete (запись помечается как удаленная, но не удаляется физически)
func (r *TimerRepo) Delete(id uint) error {
	return r.db.Delete(&models.Timer{}, id).Error
}
