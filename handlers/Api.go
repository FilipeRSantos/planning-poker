package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"liposo/planing-poker/types"
)

func (h *ApiHandler) NewRoom(ctx context.Context, user string) string {
	uid := uuid.New().String()

	//Store in a set a new room
	h.redisClient.SAdd(ctx, "rooms", uid)

	//Add the current user to the newly created room in a hash
	h.redisClient.HSet(ctx, uid, types.UserDetails{User: user, CurrentEstimate: 0})

	return uid
}

func (h *ApiHandler) HandleHealthCheck(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (h *ApiHandler) HandleNewRoom(c *fiber.Ctx) error {
	userName := c.FormValue("user")
	newRoom := h.NewRoom(c.Context(), userName)

	c.Set("hx-Location", fmt.Sprintf("/rooms/%s", newRoom))

	return c.Render("room-created", fiber.Map{
		"uid":  newRoom,
		"user": userName,
	})
}
