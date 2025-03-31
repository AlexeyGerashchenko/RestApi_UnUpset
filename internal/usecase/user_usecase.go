package usecase

import (
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
)

type UserUC struct {
	userRepo repository.UserRepository
}

func NewUserUC(userRepo repository.UserRepository) *UserUC {
	return &UserUC{userRepo}
}

func (u UserUC) Create(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) GetByID(id uint) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) GetAll() ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) Update(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) ChangePassword(id uint, oldP, newP string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) ChangeUserName(id uint, newName string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) Login(email, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}
