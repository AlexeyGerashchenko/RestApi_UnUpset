package repository

import (
	"RestApi_UnUpset/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error

	return &user, err
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "email", email).Error

	return &user, err
}

func (r *UserRepo) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *UserRepo) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepo) IsUsernameExists(username string) (bool, error) {
	var exists bool
	err := r.db.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error

	return exists, err
}
