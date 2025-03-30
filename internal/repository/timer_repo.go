package repository

import (
	"RestApi_UnUpset/internal/models"
	"gorm.io/gorm"
)

type TimerRepo struct {
	db *gorm.DB
}

func NewTimerRepo(db *gorm.DB) *TimerRepo {
	return &TimerRepo{db: db}
}

func (r *TimerRepo) Create(t *models.Timer) error {
	return r.db.Create(t).Error
}

func (r *TimerRepo) GetByID(id uint) (*models.Timer, error) {
	var t models.Timer
	err := r.db.First(&t, id).Error
	return &t, err
}

func (r *TimerRepo) Delete(id uint) error {
	return r.db.Delete(&models.Timer{}, id).Error
}
