package repository

import (
	"RestApi_UnUpset/internal/models"
	"gorm.io/gorm"
)

type StatisticsRepo struct {
	db *gorm.DB
}

func NewStatisticsRepo(db *gorm.DB) *StatisticsRepo {
	return &StatisticsRepo{db: db}
}

func (r *StatisticsRepo) Create(st *models.Statistics) error {
	return r.db.Create(st).Error
}

func (r *StatisticsRepo) GetByID(id uint) (*models.Statistics, error) {
	var st models.Statistics
	err := r.db.First(&st, id).Error
	return &st, err
}

func (r *StatisticsRepo) Update(st *models.Statistics) error {
	return r.db.Save(st).Error
}

func (r *StatisticsRepo) Delete(id uint) error {
	return r.db.Delete(&models.Statistics{}, id).Error
}
