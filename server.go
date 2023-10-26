package main

import (
	"github.com/gin-gonic/gin"
	"liposo/planing-poker/adapters"
	"liposo/planing-poker/endpoints"
)

func main() {
	_, err := adapters.CreateClient()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	endpoints.SetupPortalEndpoints(router)
	endpoints.SetupApiEndpoints(router)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
