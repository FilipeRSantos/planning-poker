package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupApiEndpoints(router *gin.Engine) {
	v1 := router.Group("/api")

	v1.GET("/healthz", handleHealthCheck)
}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
