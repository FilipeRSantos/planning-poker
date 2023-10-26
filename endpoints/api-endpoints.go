package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"liposo/planing-poker/handlers"
)

func SetupApiEndpoints(router *fiber.App, client *handlers.ApiHandler) {
	v1 := router.Group("/api")

	v1.Get("/healthz", client.HandleHealthCheck)

	roomRoutes := v1.Group("/room")
	roomRoutes.Post("/", client.HandleNewRoom)
}
