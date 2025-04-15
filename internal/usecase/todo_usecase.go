package usecase

import (
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
)

type ToDoUC struct {
	toDoRepo repository.ToDoRepository
}

func NewToDoUC(toDoRepo repository.ToDoRepository) *ToDoUC {
	return &ToDoUC{toDoRepo}
}

func (t ToDoUC) Create(toDo *models.ToDo) error {
	return t.toDoRepo.Create(toDo)
}

func (t ToDoUC) GetByID(id uint) (*models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoUC) GetByUserID(userID uint) ([]models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}

func (t ToDoUC) MarkAsDone(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (t ToDoUC) Update(toDo *models.ToDo) error {
	//TODO implement me
	panic("implement me")
}

func (t ToDoUC) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}
