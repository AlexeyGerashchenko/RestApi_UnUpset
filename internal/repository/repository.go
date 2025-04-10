// Package repository содержит интерфейсы и реализации для работы с данными в базе данных
package repository

import (
	"RestApi_UnUpset/internal/models"
	"time"

	"gorm.io/gorm"
)

// UserRepository определяет интерфейс для работы с моделью пользователя
type UserRepository interface {
	Create(user *models.User) error                 // Создание нового пользователя
	GetByID(id uint) (*models.User, error)          // Получение пользователя по ID
	GetByEmail(email string) (*models.User, error)  // Получение пользователя по email
	GetAll() ([]*models.User, error)                // Получение всех пользователей
	Update(user *models.User) error                 // Обновление данных пользователя
	IsUsernameExists(username string) (bool, error) // Проверка существования пользователя с указанным именем
	Delete(id uint) error                           // Удаление пользователя по ID
}

// ToDoRepository определяет интерфейс для работы с моделью задачи
type ToDoRepository interface {
	Create(todo *models.ToDo) error                  // Создание новой задачи
	GetByID(id uint) (*models.ToDo, error)           // Получение задачи по ID
	GetByUserID(userID uint) ([]*models.ToDo, error) // Получение всех задач пользователя
	Update(todo *models.ToDo) error                  // Обновление задачи
	Delete(id uint) error                            // Удаление задачи
}

// StatisticsRepository определяет интерфейс для работы с моделью статистики
type StatisticsRepository interface {
	Create(statistics *models.Statistics) error                 // Создание новой записи статистики
	GetByID(id uint) (*models.Statistics, error)                // Получение статистики по ID
	Update(statistics *models.Statistics) error                 // Обновление статистики
	Delete(id uint) error                                       // Удаление статистики
	GetByUserID(userID uint) (*models.Statistics, error)        // Получение статистики по ID пользователя
	IncrementCompletedTasks(userID uint) error                  // Увеличение счетчика выполненных задач
	AddFocusDuration(userID uint, duration time.Duration) error // Добавление времени фокусировки
}

// TimerRepository определяет интерфейс для работы с моделью таймера
type TimerRepository interface {
	Create(timer *models.Timer) error                 // Создание нового таймера
	GetByID(id uint) (*models.Timer, error)           // Получение таймера по ID
	GetByUserID(userID uint) ([]*models.Timer, error) // Получение всех таймеров пользователя
	Delete(id uint) error                             // Удаление таймера
}

// Repository объединяет все репозитории для удобного использования
type Repository struct {
	User       UserRepository       // Репозиторий для работы с пользователями
	ToDo       ToDoRepository       // Репозиторий для работы с задачами
	Timer      TimerRepository      // Репозиторий для работы с таймерами
	Statistics StatisticsRepository // Репозиторий для работы со статистикой
}

// NewRepository создает новый экземпляр Repository с инициализированными репозиториями
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:       NewUserRepo(db),       // Инициализация репозитория пользователей
		ToDo:       NewToDoRepo(db),       // Инициализация репозитория задач
		Timer:      NewTimerRepo(db),      // Инициализация репозитория таймеров
		Statistics: NewStatisticsRepo(db), // Инициализация репозитория статистики
	}
}
