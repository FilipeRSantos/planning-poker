package endpoints

import (
	"github.com/gin-gonic/gin"
	"liposo/planing-poker/adapters"
	"net/http"
)

func SetupPortalEndpoints(router *gin.Engine) {
	v1 := router.Group("/")

	v1.GET("/index", handleIndex)
}

func handleIndex(c *gin.Context) {
	refreshes, _ := adapters.IncAndGetRefreshed(c.Request.Context())

	c.HTML(http.StatusOK, "index.html", gin.H{
		"refreshed": refreshes,
	})
}
