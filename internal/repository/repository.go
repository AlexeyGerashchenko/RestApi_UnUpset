package repository

import (
	"RestApi_UnUpset/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type ToDoRepository interface {
	Create(todo *models.ToDo) error
	GetByID(id uint) (*models.ToDo, error)
	Update(todo *models.ToDo) error
	Delete(id uint) error
}

type StatisticsRepository interface {
	Create(statistics *models.Statistics) error
	GetByID(id uint) (*models.Statistics, error)
	Update(statistics *models.Statistics) error
	Delete(id uint) error
}

type TimerRepository interface {
	Create(timer *models.Timer) error
	GetByID(id uint) (*models.Timer, error)
	Delete(id uint) error
}

type Repository struct {
	User       UserRepository
	ToDo       ToDoRepository
	Timer      TimerRepository
	Statistics StatisticsRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:       NewUserRepo(db),
		ToDo:       NewToDoRepo(db),
		Timer:      NewTimerRepo(db),
		Statistics: NewStatisticsRepo(db),
	}
}
