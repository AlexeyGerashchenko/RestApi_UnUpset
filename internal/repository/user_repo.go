package repository

import (
	"RestApi_UnUpset/internal/models"

	"gorm.io/gorm"
)

// UserRepo представляет репозиторий для работы с моделью пользователя
type UserRepo struct {
	db *gorm.DB // Подключение к базе данных
}

// NewUserRepo создает новый экземпляр репозитория пользователей
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create добавляет нового пользователя в базу данных
func (r *UserRepo) Create(user *models.User) error {
	return r.db.Create(user).Error // Используем GORM для создания записи
}

// GetByID возвращает пользователя по его идентификатору
func (r *UserRepo) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error // Ищем первую запись с указанным ID

	return &user, err
}

// GetByEmail возвращает пользователя по его email-адресу
func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error // Используем Where для поиска по email
	return &user, err
}

// GetAll возвращает всех пользователей из базы данных
func (r *UserRepo) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error // Получаем все записи пользователей
	return users, err
}

// Update обновляет информацию о пользователе в базе данных
func (r *UserRepo) Update(user *models.User) error {
	return r.db.Save(user).Error // Save обновляет запись, если она существует
}

// Delete удаляет пользователя из базы данных по ID
// Используем Unscoped() для физического удаления, а не soft delete
func (r *UserRepo) Delete(id uint) error {
	return r.db.Unscoped().Delete(&models.User{}, id).Error
}

// IsUsernameExists проверяет, существует ли пользователь с указанным именем
func (r *UserRepo) IsUsernameExists(username string) (bool, error) {
	var count int64
	// Считаем количество пользователей с указанным именем
	err := r.db.Model(&models.User{}).
		Where("user_name = ?", username).
		Count(&count).Error
	return count > 0, err // Возвращаем true, если найден хотя бы один пользователь
}
