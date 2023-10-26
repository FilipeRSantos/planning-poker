package endpoints

import (
	"github.com/gin-gonic/gin"
	"liposo/planing-poker/handlers"
)

func SetupPortalEndpoints(router *gin.Engine, client *handlers.ApiHandler) {
	v1 := router.Group("/")

	v1.GET("/index", client.HandleIndex)
}
