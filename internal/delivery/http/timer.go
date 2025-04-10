package http

import (
	"RestApi_UnUpset/internal/delivery/dto"
	"RestApi_UnUpset/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @Summary Создание нового таймера
// @Description Создание нового таймера для текущего пользователя
// @Tags timers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body dto.CreateTimerRequest true "Данные таймера"
// @Success 201 {object} dto.Response{data=dto.TimerResponse}
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/timers [post]
func (h *Handler) createTimer(c *gin.Context) {
	session := sessions.Default(c)
	currentUserID := session.Get("user_id")
	if currentUserID == nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))
		return
	}

	var req dto.CreateTimerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	timer := &models.Timer{
		UserID:   currentUserID.(uint),
		Duration: req.Duration,
	}

	if err := h.timerUseCase.Create(timer); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	response := dto.TimerResponse{
		ID:       timer.ID,
		UserID:   timer.UserID,
		Duration: timer.Duration,
	}

	c.JSON(http.StatusCreated, dto.NewSuccessResponse(response))
}

// @Summary Получение всех таймеров пользователя
// @Description Получение списка всех таймеров текущего пользователя
// @Tags timers
// @Produce json
// @Security UserAuth
// @Success 200 {object} dto.Response{data=[]dto.TimerResponse}
// @Failure 401 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/timers [get]
func (h *Handler) getUserTimers(c *gin.Context) {
	session := sessions.Default(c)
	currentUserID := session.Get("user_id")
	if currentUserID == nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))
		return
	}

	timers, err := h.timerUseCase.GetByUserID(currentUserID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	var response []dto.TimerResponse
	for _, timer := range timers {
		response = append(response, dto.TimerResponse{
			ID:       timer.ID,
			UserID:   timer.UserID,
			Duration: timer.Duration,
		})
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}

// @Summary Получение таймера по ID
// @Description Получение конкретного таймера по его ID
// @Tags timers
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID таймера" Format(uint)
// @Success 200 {object} dto.Response{data=dto.TimerResponse}
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Router /api/timers/{id} [get]
func (h *Handler) getTimersByID(c *gin.Context) {
	timerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

	timer, err := h.timerUseCase.GetByID(uint(timerID))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if timer.UserID != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can only view your own timers"))
		return
	}

	response := dto.TimerResponse{
		ID:       timer.ID,
		UserID:   timer.UserID,
		Duration: timer.Duration,
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}

// @Summary Удаление таймера
// @Description Удаление существующего таймера (только своего)
// @Tags timers
// @Produce json
// @Security UserAuth
// @Param id path int true "ID таймера"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/timers/{id} [delete]
func (h *Handler) deleteTimer(c *gin.Context) {
	timerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

	timer, err := h.timerUseCase.GetByID(uint(timerID))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if timer.UserID != currentUserID.(uint) {
		c.JSON(http.StatusForbidden, dto.NewErrorResponse("you can only delete your own timers"))
		return
	}

	if err := h.timerUseCase.Delete(uint(timerID)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse("timer deleted successfully"))
}
