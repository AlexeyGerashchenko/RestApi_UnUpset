package usecase

import (
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
)

type TimerUC struct {
	timerRepo repository.TimerRepository
}

func NewTimerUC(timerRepo repository.TimerRepository) *TimerUC {
	return &TimerUC{timerRepo}
}

func (t TimerUC) Create(timer *models.Timer) error {
	//TODO implement me
	panic("implement me")
}

func (t TimerUC) Update(timer *models.Timer) error {
	//TODO implement me
	panic("implement me")
}

func (t TimerUC) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (t TimerUC) GetByID(id uint) (*models.Timer, error) {
	//TODO implement me
	panic("implement me")
}

func (t TimerUC) GetByUserID(userID uint) ([]models.Timer, error) {
	//TODO implement me
	panic("implement me")
}
