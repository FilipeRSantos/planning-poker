package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"liposo/planing-poker/adapters"
)

func SetupPortalEndpoints(router *fiber.App) {
	v1 := router.Group("/")

	v1.Get("/index", handleIndex)
}

func handleIndex(c *fiber.Ctx) error {
	refreshes, _ := adapters.IncAndGetRefreshed(c.Context())

	return c.Render("index", fiber.Map{
		"refreshed": refreshes,
	})
}
