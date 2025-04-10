package usecase

import (
	"RestApi_UnUpset/internal/delivery/password"
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
	"errors"
)

// UserUC реализует бизнес-логику для работы с пользователями
type UserUC struct {
	userRepo     repository.UserRepository // Репозиторий для работы с данными пользователей
	statisticsUC StatisticsUC              // Сервис для работы со статистикой
}

// NewUserUC создает новый экземпляр сервиса пользователей
func NewUserUC(userRepo repository.UserRepository, statisticsUC StatisticsUC) *UserUC {
	return &UserUC{
		userRepo:     userRepo,
		statisticsUC: statisticsUC,
	}
}

// Create создает нового пользователя и инициализирует для него статистику
func (u UserUC) Create(user *models.User) error {
	// Создаем пользователя в базе данных
	if err := u.userRepo.Create(user); err != nil {
		return err
	}

	// Инициализируем статистику для нового пользователя
	statistics := &models.Statistics{
		UserID:         user.ID,
		CompletedTasks: 0,
		FocusDuration:  0,
	}
	return u.statisticsUC.Create(statistics)
}

// GetByID возвращает пользователя по его идентификатору
func (u UserUC) GetByID(id uint) (*models.User, error) {
	return u.userRepo.GetByID(id)
}

// GetAll возвращает всех пользователей
func (u UserUC) GetAll() ([]*models.User, error) {
	return u.userRepo.GetAll()
}

// ChangePassword изменяет пароль пользователя
func (u UserUC) ChangePassword(id uint, oldP, newP string) error {
	// Получаем пользователя из базы данных
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Проверяем, соответствует ли старый пароль текущему
	if !password.CheckPassword(oldP, user.Password) {
		return errors.New("invalid old password")
	}

	// Хешируем новый пароль
	hashedPassword, err := password.HashPassword(newP)
	if err != nil {
		return err
	}

	// Обновляем пароль и сохраняем пользователя
	user.Password = hashedPassword
	return u.userRepo.Update(user)
}

// ChangeUserName изменяет имя пользователя
func (u UserUC) ChangeUserName(id uint, newName string) error {
	// Получаем пользователя из базы данных
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Проверяем, отличается ли новое имя от текущего
	if user.UserName == newName {
		return errors.New("new username should be different from the old one")
	}

	// Проверяем, не занято ли новое имя другим пользователем
	isTaken, err := u.IsUserNameTaken(newName)
	if err != nil {
		return err
	}
	if isTaken {
		return errors.New("this username is already taken")
	}

	// Обновляем имя пользователя и сохраняем
	user.UserName = newName
	return u.userRepo.Update(user)
}

// IsUserNameTaken проверяет, занято ли указанное имя пользователя
func (u UserUC) IsUserNameTaken(username string) (bool, error) {
	return u.userRepo.IsUsernameExists(username)
}

// Delete удаляет пользователя по идентификатору
func (u UserUC) Delete(id uint) error {
	return u.userRepo.Delete(id)
}

// Login аутентифицирует пользователя по email и паролю
func (u UserUC) Login(email, pw string) (*models.User, error) {
	// Получаем пользователя по email
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		// Скрываем конкретную причину ошибки для безопасности
		return nil, errors.New("invalid email or password")
	}

	// Проверяем соответствие пароля
	if !password.CheckPassword(pw, user.Password) {
		return nil, errors.New("invalid email or password2")
	}

	return user, nil
}
