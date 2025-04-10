package usecase

import (
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
)

// ToDoUC реализует бизнес-логику для работы с задачами
type ToDoUC struct {
	toDoRepo     repository.ToDoRepository // Репозиторий для работы с данными задач
	statisticsUC StatisticsUC              // Сервис для работы со статистикой
}

// NewToDoUC создает новый экземпляр сервиса задач
func NewToDoUC(toDoRepo repository.ToDoRepository, statisticsUC StatisticsUC) *ToDoUC {
	return &ToDoUC{
		toDoRepo:     toDoRepo,
		statisticsUC: statisticsUC,
	}
}

// Create добавляет новую задачу в базу данных
func (t ToDoUC) Create(toDo *models.ToDo) error {
	return t.toDoRepo.Create(toDo)
}

// GetByID возвращает задачу по её идентификатору
func (t ToDoUC) GetByID(id uint) (*models.ToDo, error) {
	return t.toDoRepo.GetByID(id)
}

// GetByUserID возвращает все активные задачи пользователя
func (t ToDoUC) GetByUserID(userID uint) ([]*models.ToDo, error) {
	return t.toDoRepo.GetByUserID(userID)
}

// MarkAsDone отмечает задачу как выполненную и обновляет статистику
func (t ToDoUC) MarkAsDone(id uint) error {
	// Получаем задачу из базы данных
	todo, err := t.toDoRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Отмечаем задачу как выполненную
	todo.Done = true
	if err := t.toDoRepo.Update(todo); err != nil {
		return err
	}

	// Увеличиваем счетчик выполненных задач в статистике
	return t.statisticsUC.IncrementCompletedTasks(todo.UserID)
}

// Update обновляет информацию о задаче
func (t ToDoUC) Update(toDo *models.ToDo) error {
	return t.toDoRepo.Update(toDo)
}

// Delete удаляет задачу по её идентификатору
func (t ToDoUC) Delete(id uint) error {
	return t.toDoRepo.Delete(id)
}
