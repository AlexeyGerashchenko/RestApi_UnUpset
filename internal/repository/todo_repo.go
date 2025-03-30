package repository

import (
	"RestApi_UnUpset/internal/models"
	"gorm.io/gorm"
)

type ToDoRepo struct {
	db *gorm.DB
}

func NewToDoRepo(db *gorm.DB) *ToDoRepo {
	return &ToDoRepo{db: db}
}

func (r *ToDoRepo) Create(toDo *models.ToDo) error {
	return r.db.Create(toDo).Error
}

func (r *ToDoRepo) GetByID(id uint) (*models.ToDo, error) {
	var toDo models.ToDo
	err := r.db.First(&toDo, id).Error
	return &toDo, err
}

func (r *ToDoRepo) Update(toDo *models.ToDo) error {
	return r.db.Save(toDo).Error
}

func (r *ToDoRepo) Delete(id uint) error {
	return r.db.Delete(&models.ToDo{}, id).Error
}
