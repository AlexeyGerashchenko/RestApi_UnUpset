package http

import (
	"RestApi_UnUpset/internal/delivery/middleware"
	"RestApi_UnUpset/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUseCase       usecase.UserUseCase
	todoUseCase       usecase.ToDoUseCase
	timerUseCase      usecase.TimerUseCase
	statisticsUseCase usecase.StatisticsUseCase
}

func NewHandler(useCase *usecase.UseCase) *Handler {
	return &Handler{
		userUseCase:       useCase.User,
		todoUseCase:       useCase.ToDo,
		timerUseCase:      useCase.Timer,
		statisticsUseCase: useCase.Statistics,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/register", h.register)
	}

	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		user := protected.Group("/user")
		{
			user.GET("/", h.getAllUsers)
			user.GET("/:id", h.getByID)
			user.PUT("/:id", h.updateUser)
			user.PATCH("/password/:id", h.changePassword)
			user.PATCH("/username/:id", h.changeUserName)
			user.DELETE("/:id", h.deleteUser)
		}

		todos := protected.Group("/todos")
		{
			todos.POST("/", h.createToDo)
			todos.GET("/", h.getUserToDos)
			todos.GET("/:id", h.getToDoByID)
			todos.PUT("/:id", h.updateToDo)
			todos.PATCH("/:id/done", h.markToDoDone)
			todos.DELETE("/:id", h.deleteToDo)
		}

		statistics := protected.Group("/statistics")
		{
			statistics.POST("/", h.createStatistics)
			statistics.GET("/", h.getUserStatistics)
			statistics.GET("/:id", h.getStatisticsByID)
			statistics.GET("/filter", h.getWithFilter)
			statistics.PUT("/:id", h.updateStatistics)
			statistics.DELETE("/:id", h.deleteStatistics)
		}

		timers := protected.Group("/timers")
		{
			timers.POST("/", h.createTimer)
			timers.GET("/", h.getUserTimers)
			timers.GET("/:id", h.getTimersByID)
			timers.PUT("/:id", h.updateTimer)
			timers.DELETE("/:id", h.deleteTimer)
		}
	}

	return router
}
