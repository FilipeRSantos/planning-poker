package endpoints

import (
	"github.com/gin-gonic/gin"
	"liposo/planing-poker/handlers"
)

func SetupApiEndpoints(router *gin.Engine, client *handlers.ApiHandler) {
	v1 := router.Group("/api")

	v1.GET("/healthz", client.HandleHealthCheck)

	roomRoutes := v1.Group("/room")
	roomRoutes.POST("/", client.HandleNewRoom)
}
