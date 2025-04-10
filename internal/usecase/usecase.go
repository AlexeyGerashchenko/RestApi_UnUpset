// Package usecase содержит бизнес-логику приложения, обрабатывающую данные
package usecase

import (
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
	"time"
)

// UserUseCase определяет интерфейс для работы с бизнес-логикой пользователей
type UserUseCase interface {
	Create(user *models.User) error                     // Создание нового пользователя
	GetByID(id uint) (*models.User, error)              // Получение пользователя по ID
	GetAll() ([]*models.User, error)                    // Получение всех пользователей
	ChangePassword(id uint, oldP, newP string) error    // Изменение пароля пользователя
	IsUserNameTaken(username string) (bool, error)      // Проверка занятости имени пользователя
	ChangeUserName(id uint, newName string) error       // Изменение имени пользователя
	Delete(id uint) error                               // Удаление пользователя
	Login(email, password string) (*models.User, error) // Аутентификация пользователя
}

// ToDoUseCase определяет интерфейс для работы с бизнес-логикой задач
type ToDoUseCase interface {
	Create(toDo *models.ToDo) error                  // Создание новой задачи
	Update(toDo *models.ToDo) error                  // Обновление задачи
	Delete(id uint) error                            // Удаление задачи
	GetByID(id uint) (*models.ToDo, error)           // Получение задачи по ID
	GetByUserID(userID uint) ([]*models.ToDo, error) // Получение всех задач пользователя
	MarkAsDone(id uint) error                        // Отметка задачи как выполненной
}

// StatisticsUseCase определяет интерфейс для работы с бизнес-логикой статистики
type StatisticsUseCase interface {
	Create(statistics *models.Statistics) error                 // Создание новой записи статистики
	Update(statistics *models.Statistics) error                 // Обновление статистики
	Delete(id uint) error                                       // Удаление статистики
	GetByID(id uint) (*models.Statistics, error)                // Получение статистики по ID
	GetByUserID(userID uint) (*models.Statistics, error)        // Получение статистики пользователя
	IncrementCompletedTasks(userID uint) error                  // Увеличение счетчика выполненных задач
	AddFocusDuration(userID uint, duration time.Duration) error // Добавление времени фокусировки
}

// TimerUseCase определяет интерфейс для работы с бизнес-логикой таймеров
type TimerUseCase interface {
	Create(timer *models.Timer) error                 // Создание нового таймера
	Delete(id uint) error                             // Удаление таймера
	GetByID(id uint) (*models.Timer, error)           // Получение таймера по ID
	GetByUserID(userID uint) ([]*models.Timer, error) // Получение всех таймеров пользователя
}

// UseCase объединяет все сервисы бизнес-логики для удобного использования
type UseCase struct {
	User       UserUseCase       // Сервис для работы с пользователями
	ToDo       ToDoUseCase       // Сервис для работы с задачами
	Timer      TimerUseCase      // Сервис для работы с таймерами
	Statistics StatisticsUseCase // Сервис для работы со статистикой
}

// NewUseCase создает новый экземпляр UseCase с инициализированными сервисами
func NewUseCase(rep *repository.Repository) *UseCase {
	// Сначала создаем сервис статистики, так как он используется другими сервисами
	statisticsUC := NewStatisticsUC(rep.Statistics)
	return &UseCase{
		User:       NewUserUC(rep.User, *statisticsUC),   // Инициализация сервиса пользователей
		ToDo:       NewToDoUC(rep.ToDo, *statisticsUC),   // Инициализация сервиса задач
		Timer:      NewTimerUC(rep.Timer, *statisticsUC), // Инициализация сервиса таймеров
		Statistics: statisticsUC,                         // Инициализация сервиса статистики
	}
}
