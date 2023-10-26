package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *ApiHandler) IncAndGetRefreshed(ctx context.Context) (int64, error) {
	return h.redisClient.Incr(ctx, "refreshes").Result()
}

func (h *ApiHandler) HandleIndex(c *gin.Context) {
	refreshes, _ := h.IncAndGetRefreshed(c.Request.Context())

	c.HTML(http.StatusOK, "index.html", gin.H{
		"refreshed": refreshes,
	})
}
