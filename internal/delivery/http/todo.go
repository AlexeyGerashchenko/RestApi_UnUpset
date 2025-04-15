package http

import (
	"RestApi_UnUpset/internal/delivery/dto"
	"RestApi_UnUpset/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func (h *Handler) getUserToDos(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
func (h *Handler) getToDoByID(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
func (h *Handler) updateToDo(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) markToDoDone(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) deleteToDo(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
