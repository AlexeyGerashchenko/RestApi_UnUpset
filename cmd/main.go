package main

import (
	"RestApi_UnUpset/docs"
	"RestApi_UnUpset/internal/delivery/http"
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
	"RestApi_UnUpset/internal/usecase"
	"log"
	h "net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           UnUpset API
// @version         1.0
// @description     API для приложения UnUpset
// @host      localhost:8080

// @securityDefinitions.apikey ApiKeyAuth
// @in cookie
// @name mysession
// @description Для тестирования в Swagger UI (на макбуке логин работает только через хром, в сафари протестировать не получается)
func main() {
	// Получаем строку подключения к базе данных из переменной окружения
	dsn := os.Getenv("DB_DSN")
	// Инициализируем соединение с PostgreSQL базой данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Логируем и завершаем программу в случае ошибки подключения
		log.Fatal("Failed to connect to database:", err)
	}
	// Автоматическая миграция моделей в базу данных
	// Создаем таблицы для всех наших моделей, если они еще не существуют
	if err := db.AutoMigrate(
		&models.User{},
		&models.ToDo{},
		&models.Timer{},
		&models.Statistics{},
	); err != nil {
		// Логируем и завершаем программу в случае ошибки миграции
		log.Fatal("Failed to migrate database:", err)
	}

	// Инициализируем репозиторий для работы с базой данных
	repo := repository.NewRepository(db)
	// Инициализируем слой бизнес-логики, передавая ему репозиторий
	useCase := usecase.NewUseCase(repo)
	// Инициализируем обработчики HTTP-запросов, передавая им слой бизнес-логики
	handler := http.NewHandler(useCase)

	// Создаем новый экземпляр Gin-роутера с настройками по умолчанию
	router := gin.Default()

	// Настраиваем куки-хранилище для сессий пользователей
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	// Настраиваем параметры безопасности для куки
	store.Options(sessions.Options{
		HttpOnly: true,              // Куки доступны только через HTTP, не через JavaScript
		Secure:   true,              // Куки передаются только по HTTPS
		SameSite: h.SameSiteLaxMode, // Защита от CSRF-атак
		MaxAge:   86400 * 7,         // Срок жизни куки - 7 дней в секундах
		Path:     "/",               // Куки доступны на всех путях приложения
	})

	// Добавляем middleware для работы с сессиями
	router.Use(sessions.Sessions("mysession", store))

	// Инициализируем маршруты для API
	router = handler.InitRoutes(router)

	// Настройка Swagger документации
	docs.SwaggerInfo.Title = "UnUpset API"
	docs.SwaggerInfo.Description = "API для приложения UnUpset"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Добавляем маршрут для Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запускаем HTTP-сервер на порту 8080
	if err := router.Run(":8080"); err != nil {
		// Логируем и завершаем программу в случае ошибки запуска сервера
		log.Fatal("Failed to start server:", err)
	}
}
