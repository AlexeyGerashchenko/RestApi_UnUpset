package repository

import (
	"RestApi_UnUpset/internal/models"
	"time"

	"gorm.io/gorm"
)

// StatisticsRepo представляет репозиторий для работы с моделью статистики
type StatisticsRepo struct {
	db *gorm.DB // Подключение к базе данных
}

// NewStatisticsRepo создает новый экземпляр репозитория статистики
func NewStatisticsRepo(db *gorm.DB) *StatisticsRepo {
	return &StatisticsRepo{db: db}
}

// GetByUserID возвращает статистику конкретного пользователя
func (r *StatisticsRepo) GetByUserID(userID uint) (*models.Statistics, error) {
	var statistics models.Statistics
	// Ищем первую запись статистики для указанного пользователя
	err := r.db.Where("user_id = ?", userID).First(&statistics).Error
	if err != nil {
		return nil, err
	}
	return &statistics, nil
}

// IncrementCompletedTasks увеличивает счетчик завершенных задач пользователя на 1
func (r *StatisticsRepo) IncrementCompletedTasks(userID uint) error {
	// Используем SQL-выражение для инкремента без необходимости получать всю запись
	return r.db.Model(&models.Statistics{}).
		Where("user_id = ?", userID).
		UpdateColumn("completed_tasks", gorm.Expr("completed_tasks + ?", 1)).
		Error
}

// AddFocusDuration увеличивает общее время фокусировки пользователя на указанную продолжительность
func (r *StatisticsRepo) AddFocusDuration(userID uint, duration time.Duration) error {
	// Используем SQL-выражение для добавления времени без необходимости получать всю запись
	return r.db.Model(&models.Statistics{}).
		Where("user_id = ?", userID).
		UpdateColumn("focus_duration", gorm.Expr("focus_duration + ?", duration)).
		Error
}

// Create добавляет новую запись статистики в базу данных
func (r *StatisticsRepo) Create(st *models.Statistics) error {
	return r.db.Create(st).Error // Используем GORM для создания записи
}

// GetByID возвращает статистику по её идентификатору
func (r *StatisticsRepo) GetByID(id uint) (*models.Statistics, error) {
	var st models.Statistics
	err := r.db.First(&st, id).Error // Ищем первую запись с указанным ID
	return &st, err
}

// Update обновляет информацию о статистике в базе данных
func (r *StatisticsRepo) Update(st *models.Statistics) error {
	return r.db.Save(st).Error // Save обновляет запись, если она существует
}

// Delete удаляет статистику из базы данных по ID
// Используется soft delete (запись помечается как удаленная, но не удаляется физически)
func (r *StatisticsRepo) Delete(id uint) error {
	return r.db.Delete(&models.Statistics{}, id).Error
}
