package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthCredentials структура для хранения данных аутентификации
type AuthCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool `json:"success"`
}

func RouteAuth(r *gin.Engine) {
	authGroup := r.Group("/auth")

	authGroup.POST("/jwt", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"jwt": "your_token_here", // Здесь должна быть генерация и отправка настоящего JWT
		})
	})

	authGroup.POST("/credCheck", func(c *gin.Context) {
		var creds AuthCredentials

		if err := c.BindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		success := creds.Username == "sum" && creds.Password == "jwt"

		c.JSON(http.StatusOK, gin.H{"success": success})
	})
}
