package http

import (
	"RestApi_UnUpset/internal/delivery/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Получение статистики пользователя
// @Description Получение статистики о выполненных задачах и времени фокусировки
// @Tags statistics
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} dto.Response{data=dto.StatisticsResponse}
// @Failure 401 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/statistics [get]
func (h *Handler) getUserStatistics(c *gin.Context) {
	userID := c.GetUint("user_id")

	statistics, err := h.statisticsUseCase.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Failed to get statistics"))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(dto.StatisticsResponse{
		CompletedTasks: statistics.CompletedTasks,
		FocusDuration:  statistics.FocusDuration.String(),
	}))
}
