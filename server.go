package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"onlineUsersCount": 1,
		})
	})

	err := router.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}
