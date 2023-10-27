package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"liposo/planing-poker/handlers"
)

func SetupPortalEndpoints(router *fiber.App, client *handlers.ApiHandler) {
	v1 := router.Group("/")

	v1.Get("/", client.HandleIndex)
	v1.Get("/rooms/:uid", client.HandleGetRoomByUid)
}
