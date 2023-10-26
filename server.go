package main

import (
	"github.com/gin-gonic/gin"
	"liposo/planing-poker/endpoints"
	"liposo/planing-poker/handlers"
)

func main() {
	liposoClient, err := handlers.CreateClient()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	endpoints.SetupPortalEndpoints(router, liposoClient)
	endpoints.SetupApiEndpoints(router, liposoClient)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
