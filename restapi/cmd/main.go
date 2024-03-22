package main

import (
	"net/http"

	routes "github.com/XRS0/Forms/restapi/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Сервирование статических файлов для фронтенда
	r.Static("/styles", "/frontend/styles/")
	r.Static("/scripts", "/frontend/scripts/")
	r.LoadHTMLGlob("/frontend/templates/*.html")

	// r.Static("/styles", "../../frontend/styles/")
	// r.Static("/scripts", "../../frontend/scripts/")
	// r.LoadHTMLGlob("../../frontend/templates/*.html")

	// Корневой маршрут, который загружает главную HTML страницу
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	routes.RouteAuth(r)
	r.Run(":8080")
}
