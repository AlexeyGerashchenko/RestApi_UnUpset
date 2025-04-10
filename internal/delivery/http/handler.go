// Package http содержит обработчики HTTP-запросов для API
package http

import (
	"RestApi_UnUpset/internal/delivery/middleware"
	"RestApi_UnUpset/internal/usecase"

	"github.com/gin-gonic/gin"
)

// Handler структура для обработки HTTP-запросов
type Handler struct {
	userUseCase       usecase.UserUseCase       // Сервис бизнес-логики для пользователей
	todoUseCase       usecase.ToDoUseCase       // Сервис бизнес-логики для задач
	timerUseCase      usecase.TimerUseCase      // Сервис бизнес-логики для таймеров
	statisticsUseCase usecase.StatisticsUseCase // Сервис бизнес-логики для статистики
}

// NewHandler создает новый экземпляр обработчика с нужными сервисами
func NewHandler(useCase *usecase.UseCase) *Handler {
	return &Handler{
		userUseCase:       useCase.User,       // Инициализация сервиса пользователей
		todoUseCase:       useCase.ToDo,       // Инициализация сервиса задач
		timerUseCase:      useCase.Timer,      // Инициализация сервиса таймеров
		statisticsUseCase: useCase.Statistics, // Инициализация сервиса статистики
	}
}

// InitRoutes инициализирует маршруты для API
func (h *Handler) InitRoutes(router *gin.Engine) *gin.Engine {

	// Группа маршрутов для авторизации (без аутентификации)
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)       // Маршрут для авторизации
		auth.POST("/register", h.register) // Маршрут для регистрации
	}

	// Группа защищенных маршрутов (требуется аутентификация)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware()) // Добавляем middleware для проверки аутентификации
	{
		// Маршруты для работы с пользователями
		user := protected.Group("/user")
		{
			user.GET("/", h.getAllUsers)                  // Получить всех пользователей
			user.GET("/:id", h.getByID)                   // Получить пользователя по ID
			user.PATCH("/password/:id", h.changePassword) // Изменить пароль
			user.PATCH("/username/:id", h.changeUserName) // Изменить имя пользователя
			user.DELETE("/:id", h.deleteUser)             // Удалить пользователя
		}

		// Маршруты для работы с задачами
		todos := protected.Group("/todos")
		{
			todos.POST("/", h.createToDo)            // Создать новую задачу
			todos.GET("/", h.getUserToDos)           // Получить все задачи пользователя
			todos.GET("/:id", h.getToDoByID)         // Получить задачу по ID
			todos.PUT("/:id", h.updateToDo)          // Обновить задачу
			todos.PATCH("/:id/done", h.markToDoDone) // Отметить задачу как выполненную
			todos.DELETE("/:id", h.deleteToDo)       // Удалить задачу
		}

		// Маршруты для работы со статистикой
		statistics := protected.Group("/statistics")
		{
			statistics.GET("/", h.getUserStatistics) // Получить статистику пользователя
		}

		// Маршруты для работы с таймерами
		timers := protected.Group("/timers")
		{
			timers.POST("/", h.createTimer)      // Создать новый таймер
			timers.GET("/", h.getUserTimers)     // Получить все таймеры пользователя
			timers.GET("/:id", h.getTimersByID)  // Получить таймер по ID
			timers.DELETE("/:id", h.deleteTimer) // Удалить таймер
		}
	}

	return router
}
