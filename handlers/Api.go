package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"liposo/planing-poker/types"
	"net/http"
)

func (h *ApiHandler) NewRoom(ctx context.Context, user string) string {
	uid := uuid.New().String()

	//Store in a set a new room
	h.redisClient.SAdd(ctx, "rooms", uid)

	//Add the current user to the newly created room in a hash
	h.redisClient.HSet(ctx, uid, types.UserDetails{User: user, CurrentEstimate: 0})

	return uid
}

func (h *ApiHandler) HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (h *ApiHandler) HandleNewRoom(c *gin.Context) {
	userName := c.PostForm("user")
	newRoom := h.NewRoom(c.Request.Context(), userName)

	c.HTML(http.StatusCreated, "room-created.html", gin.H{
		"uid":  newRoom,
		"user": userName,
	})
}
