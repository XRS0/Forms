package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
			return
		}

		tokenString := authHeader[len(BearerSchema):]
		// Здесь должна быть ваша логика проверки JWT токена
		// Например, можно распарсить токен, проверить его подпись и срок действия
		// Если токен невалиден, верните ошибку

		if !isTokenValid(tokenString) { // Представьте, что это ваша функция проверки токена
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API token"})
			return
		}

		c.Next() // Продолжить выполнение, если токен валиден
	}
}

func isTokenValid(token string) bool {

	return true
}
