package endpoints

import (
	"github.com/gin-gonic/gin"
	"liposo/planing-poker/adapters"
	"net/http"
)

func SetupApiEndpoints(router *gin.Engine) {
	v1 := router.Group("/api")

	v1.GET("/healthz", handleHealthCheck)

	roomRoutes := v1.Group("/room")
	roomRoutes.POST("/", handleNewRoom)
}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func handleNewRoom(c *gin.Context) {
	userName := c.PostForm("user")
	newRoom := adapters.NewRoom(c.Request.Context(), userName)

	c.HTML(http.StatusCreated, "room-created.html", gin.H{
		"uid":  newRoom,
		"user": userName,
	})
}
