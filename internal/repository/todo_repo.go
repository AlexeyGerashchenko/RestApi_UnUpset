package repository

import (
	"RestApi_UnUpset/internal/models"

	"gorm.io/gorm"
)

// ToDoRepo представляет репозиторий для работы с моделью задач
type ToDoRepo struct {
	db *gorm.DB // Подключение к базе данных
}

// NewToDoRepo создает новый экземпляр репозитория задач
func NewToDoRepo(db *gorm.DB) *ToDoRepo {
	return &ToDoRepo{db: db}
}

// Create добавляет новую задачу в базу данных
func (r *ToDoRepo) Create(toDo *models.ToDo) error {
	return r.db.Create(toDo).Error // Используем GORM для создания записи
}

// GetByID возвращает задачу по её идентификатору
func (r *ToDoRepo) GetByID(id uint) (*models.ToDo, error) {
	var toDo models.ToDo
	err := r.db.First(&toDo, id).Error // Ищем первую запись с указанным ID
	return &toDo, err
}

// GetByUserID возвращает все активные (невыполненные) задачи конкретного пользователя
func (r *ToDoRepo) GetByUserID(userID uint) ([]*models.ToDo, error) {
	var todos []*models.ToDo
	// Фильтруем по ID пользователя и статусу выполнения (только невыполненные)
	err := r.db.Where("user_id = ? AND done = ?", userID, false).Find(&todos).Error
	return todos, err
}

// Update обновляет информацию о задаче в базе данных
func (r *ToDoRepo) Update(toDo *models.ToDo) error {
	return r.db.Save(toDo).Error // Save обновляет запись, если она существует
}

// Delete удаляет задачу из базы данных по ID
// Используется soft delete (запись помечается как удаленная, но не удаляется физически)
func (r *ToDoRepo) Delete(id uint) error {
	return r.db.Delete(&models.ToDo{}, id).Error
}
