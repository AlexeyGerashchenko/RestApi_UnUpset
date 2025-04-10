// Package middleware содержит промежуточные обработчики для HTTP-запросов
package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware создает middleware для проверки аутентификации пользователя
// Проверяет наличие идентификатора пользователя в сессии
// Если пользователь не аутентифицирован, запрос прерывается с кодом 401
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем сессию из контекста запроса
		session := sessions.Default(c)
		// Получаем ID пользователя из сессии
		userID := session.Get("user_id")

		// Если ID пользователя отсутствует, значит пользователь не аутентифицирован
		if userID == nil {
			// Прерываем обработку запроса с ошибкой 401 Unauthorized
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		// Добавляем ID пользователя в контекст для использования в обработчиках
		c.Set("user_id", userID)
		// Продолжаем обработку запроса
		c.Next()
	}
}
