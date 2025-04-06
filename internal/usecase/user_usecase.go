package usecase

import (
	"RestApi_UnUpset/internal/delivery/password"
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
	"errors"
)

type UserUC struct {
	userRepo repository.UserRepository
}

func NewUserUC(userRepo repository.UserRepository) *UserUC {
	return &UserUC{userRepo}
}

func (u UserUC) Create(user *models.User) error {
	hashedPassword, err := password.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return u.userRepo.Create(user)
}

func (u UserUC) GetByID(id uint) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) GetAll() ([]*models.User, error) {
	return u.userRepo.GetAll()
}

func (u UserUC) Update(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) ChangePassword(id uint, oldP, newP string) error {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	if !password.CheckPassword(oldP, user.Password) {
		return errors.New("invalid old password")
	}

	hashedPassword, err := password.HashPassword(newP)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return u.userRepo.Update(user)
}

func (u UserUC) ChangeUserName(id uint, newName string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) IsUserNameTaken(username string) (bool, error) {
	return u.userRepo.IsUsernameExists(username)
}

func (u UserUC) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUC) Login(email, pw string) (*models.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if !password.CheckPassword(pw, user.Password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}
