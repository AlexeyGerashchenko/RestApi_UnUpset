package http

import (
	"RestApi_UnUpset/internal/delivery/dto"
	"RestApi_UnUpset/internal/delivery/password"
	"RestApi_UnUpset/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @Summary Вход пользователя
// @Description Аутентификация пользователя по email и паролю
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest true "Данные для входа"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	user, err := h.userUseCase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.NewSuccessResponse(nil))
}

// @Summary Регистрация пользователя
// @Description Создание нового пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterUserRequest true "Данные для регистрации"
// @Success 201 {object} dto.Response{data=dto.UserResponse}
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /auth/register [post]
func (h *Handler) register(c *gin.Context) {
	var req dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	user := &models.User{
		UserName: req.UserName,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := h.userUseCase.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	response := dto.UserResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, dto.NewSuccessResponse(response))
}

// @Summary Получение списка всех пользователей
// @Description Получение списка всех зарегистрированных пользователей
// @Tags users
// @Produce json
// @Success 200 {object} dto.Response{data=[]dto.UserResponse}
// @Failure 500 {object} dto.Response
// @Router /api/user [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.userUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	var response []dto.UserResponse
	for _, user := range users {
		response = append(response, dto.UserResponse{
			ID:        user.ID,
			UserName:  user.UserName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}

// @Summary Получение информации о пользователе
// @Description Получение информации о конкретном пользователе по ID
// @Tags users
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID пользователя" Format(uint)
// @Success 200 {object} dto.Response{data=dto.UserResponse}
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Router /api/user/{id} [get]
func (h *Handler) getByID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	session := sessions.Default(c)
	currentUserID := session.Get("user_id")
	if currentUserID == nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))
		return
	}

	if uint(userID) != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can get only your own information"))
		return
	}

	user, err := h.userUseCase.GetByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	response := dto.UserResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}

// @Summary Изменение пароля
// @Description Изменение пароля пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param input body dto.ChangePasswordRequest true "Данные для смены пароля"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Router /api/user/password/{id} [patch]
func (h *Handler) changePassword(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	session := sessions.Default(c)
	currentUserID := session.Get("user_id")
	if currentUserID == nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))
		return
	}

	if uint(userID) != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can only change your own password"))
		return
	}

	var req dto.ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	err = h.userUseCase.ChangePassword(uint(userID), req.OldPassword, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse("password changed successfully"))
}

// @Summary Изменение имени пользователя
// @Description Изменение имени пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param input body dto.ChangeUsernameRequest true "Новое имя пользователя"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Router /api/user/username/{id} [patch]
func (h *Handler) changeUserName(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	session := sessions.Default(c)
	currentUserID := session.Get("user_id")
	if currentUserID == nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))
		return
	}

	if uint(userID) != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can only change your own username"))
		return
	}

	var req dto.ChangeUsernameRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	err = h.userUseCase.ChangeUserName(uint(userID), req.NewUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse("username changed successfully"))
}

// @Summary Удаление пользователя
// @Description Удаление пользователя по ID
// @Tags users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Router /api/user/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	session := sessions.Default(c)
	currentUserID := session.Get("user_id")
	if currentUserID == nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))
		return
	}

	if uint(userID) != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can only delete your own account"))
		return
	}

	err = h.userUseCase.Delete(uint(userID))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse("account deleted successfully"))
}
