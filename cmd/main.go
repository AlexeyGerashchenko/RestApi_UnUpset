package main

import (
	"RestApi_UnUpset/internal/delivery/http"
	"RestApi_UnUpset/internal/models"
	"RestApi_UnUpset/internal/repository"
	"RestApi_UnUpset/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	h "net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err := db.AutoMigrate(
		&models.User{},
		&models.ToDo{},
		&models.Timer{},
		&models.Statistics{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	repo := repository.NewRepository(db)
	useCase := usecase.NewUseCase(repo)
	handler := http.NewHandler(useCase)
	router := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	store.Options(sessions.Options{
		HttpOnly: true,
		Secure:   true,
		SameSite: h.SameSiteLaxMode,
		MaxAge:   86400 * 7,
	})
	router.Use(sessions.Sessions("mysession", store))
	router = handler.InitRoutes(router)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
