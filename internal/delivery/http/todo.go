package http

import (
	"RestApi_UnUpset/internal/delivery/dto"
	"RestApi_UnUpset/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @Summary Создание новой задачи
// @Description Создание новой задачи для текущего пользователя
// @Tags todos
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body dto.CreateToDoRequest true "Данные задачи"
// @Success 201 {object} dto.Response{data=dto.ToDorResponse}
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/todos [post]
func (h *Handler) createToDo(c *gin.Context) {
	session := sessions.Default(c)
	currentUserID := session.Get("user_id")
	if currentUserID == nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))
		return
	}

	var req dto.CreateToDoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	todo := &models.ToDo{
		UserID: currentUserID.(uint),
		Text:   req.Text,
	}

	if err := h.todoUseCase.Create(todo); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	response := dto.ToDorResponse{
		ID:     todo.ID,
		UserID: todo.UserID,
		Text:   todo.Text,
		Done:   todo.Done,
	}

	c.JSON(http.StatusCreated, dto.NewSuccessResponse(response))

}

// @Summary Получение всех задач пользователя (не выполненных)
// @Description Получение списка всех задач текущего пользователя (не выполненных)
// @Tags todos
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} dto.Response{data=[]dto.ToDorResponse}
// @Failure 401 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/todos [get]
func (h *Handler) getUserToDos(c *gin.Context) {
	session := sessions.Default(c)
	currentUserID := session.Get("user_id")
	if currentUserID == nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))
		return
	}

	todos, err := h.todoUseCase.GetByUserID(currentUserID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	var response []dto.ToDorResponse
	for _, todo := range todos {
		response = append(response, dto.ToDorResponse{
			ID:     todo.ID,
			UserID: todo.UserID,
			Text:   todo.Text,
			Done:   todo.Done,
		})
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}

// @Summary Получение задачи по ID
// @Description Получение конкретной задачи по её ID
// @Tags todos
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID задачи" Format(uint)
// @Success 200 {object} dto.Response{data=dto.ToDorResponse}
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Router /api/todos/{id} [get]
func (h *Handler) getToDoByID(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

	todo, err := h.todoUseCase.GetByID(uint(todoID))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if todo.UserID != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can only view your own todos"))
		return
	}

	response := dto.ToDorResponse{
		ID:     todo.ID,
		UserID: todo.UserID,
		Text:   todo.Text,
		Done:   todo.Done,
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}

// @Summary Обновление задачи
// @Description Обновление текста существующей задачи
// @Tags todos
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID задачи"
// @Param input body dto.CreateToDoRequest true "Новые данные задачи"
// @Success 200 {object} dto.Response{data=dto.ToDorResponse}
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/todos/{id} [put]
func (h *Handler) updateToDo(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

	var req dto.CreateToDoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	todo, err := h.todoUseCase.GetByID(uint(todoID))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if todo.UserID != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can only update your own todos"))
		return
	}

	todo.Text = req.Text

	if err := h.todoUseCase.Update(todo); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	response := dto.ToDorResponse{
		ID:     todo.ID,
		UserID: todo.UserID,
		Text:   todo.Text,
		Done:   todo.Done,
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}

// @Summary Отметить задачу как выполненную
// @Description Отметить существующую задачу как выполненную
// @Tags todos
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID задачи"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/todos/{id}/done [patch]
func (h *Handler) markToDoDone(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

	todo, err := h.todoUseCase.GetByID(uint(todoID))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if todo.UserID != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can only mark your own todos as done"))
		return
	}

	if err := h.todoUseCase.MarkAsDone(uint(todoID)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse("todo marked as done"))
}

// @Summary Удаление задачи
// @Description Удаление существующей задачи
// @Tags todos
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID задачи"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/todos/{id} [delete]
func (h *Handler) deleteToDo(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

	todo, err := h.todoUseCase.GetByID(uint(todoID))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if todo.UserID != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can only delete your own todos"))
		return
	}

	if err := h.todoUseCase.Delete(uint(todoID)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse("todo deleted successfully"))
}
