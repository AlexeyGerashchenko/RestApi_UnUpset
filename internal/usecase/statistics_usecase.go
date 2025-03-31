package usecase

import (
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
	"time"
)

type StatisticsUC struct {
	statisticsRepo repository.StatisticsRepository
}

func NewStatisticsUC(statisticsRepo repository.StatisticsRepository) *StatisticsUC {
	return &StatisticsUC{statisticsRepo}
}

func (s StatisticsUC) Create(statistics *models.Statistics) error {
	//TODO implement me
	panic("implement me")
}

func (s StatisticsUC) Update(statistics *models.Statistics) error {
	//TODO implement me
	panic("implement me")
}

func (s StatisticsUC) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (s StatisticsUC) GetByID(id uint) (*models.Statistics, error) {
	//TODO implement me
	panic("implement me")
}

func (s StatisticsUC) GetByUserID(userID uint) ([]models.Statistics, error) {
	//TODO implement me
	panic("implement me")
}

func (s StatisticsUC) FilterByDates(userID uint, startTime, endTime time.Time) ([]models.Statistics, error) {
	//TODO implement me
	panic("implement me")
}
