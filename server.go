package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
)

var client *redis.Client

func main() {
	opt, err := redis.ParseURL(os.Getenv("REDIS_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	client = redis.NewClient(opt)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/index", func(c *gin.Context) {
		refreshes := incAndGetRefreshed(c.Request.Context())

		c.HTML(http.StatusOK, "index.html", gin.H{
			"refreshed": refreshes,
		})
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	err = router.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

func incAndGetRefreshed(ctx context.Context) int64 {
	val, err := client.Incr(ctx, "refreshes").Result()
	if err != nil {
		panic(err)
	}

	return val
}
