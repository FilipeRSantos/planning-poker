package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"liposo/planing-poker/adapters"
)

func SetupApiEndpoints(router *fiber.App) {
	v1 := router.Group("/api")

	v1.Get("/healthz", handleHealthCheck)

	roomRoutes := v1.Group("/room")
	roomRoutes.Post("/", handleNewRoom)
}

func handleHealthCheck(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func handleNewRoom(c *fiber.Ctx) error {
	userName := c.FormValue("user")
	newRoom := adapters.NewRoom(c.Context(), userName)

	return c.Render("room-created", fiber.Map{
		"uid":  newRoom,
		"user": userName,
	})
}
