package usecase

import (
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
	"time"
)

type UserUseCase interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(user *models.User) error
	ChangePassword(id uint, oldP, newP string) error
	IsUserNameTaken(username string) (bool, error)
	ChangeUserName(id uint, newName string) error
	Delete(id uint) error
	Login(email, password string) (*models.User, error)
}

type ToDoUseCase interface {
	Create(toDo *models.ToDo) error
	Update(toDo *models.ToDo) error
	Delete(id uint) error
	GetByID(id uint) (*models.ToDo, error)
	GetByUserID(userID uint) ([]models.ToDo, error)
	MarkAsDone(id uint) error
}

type StatisticsUseCase interface {
	Create(statistics *models.Statistics) error
	Update(statistics *models.Statistics) error
	Delete(id uint) error
	GetByID(id uint) (*models.Statistics, error)
	GetByUserID(userID uint) ([]models.Statistics, error)
	FilterByDates(userID uint, startTime, endTime time.Time) ([]models.Statistics, error)
}

type TimerUseCase interface {
	Create(timer *models.Timer) error
	Update(timer *models.Timer) error
	Delete(id uint) error
	GetByID(id uint) (*models.Timer, error)
	GetByUserID(userID uint) ([]models.Timer, error)
}

type UseCase struct {
	User       UserUseCase
	ToDo       ToDoUseCase
	Timer      TimerUseCase
	Statistics StatisticsUseCase
}

func NewUseCase(rep *repository.Repository) *UseCase {
	return &UseCase{
		User:       NewUserUC(rep.User),
		ToDo:       NewToDoUC(rep.ToDo),
		Timer:      NewTimerUC(rep.Timer),
		Statistics: NewStatisticsUC(rep.Statistics),
	}
}
