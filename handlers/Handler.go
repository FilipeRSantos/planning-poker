package handlers

import (
	"github.com/redis/go-redis/v9"
	"os"
)

type ApiHandler struct {
	redisClient *redis.Client
}

func CreateClient() (*ApiHandler, error) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}

	handler := &ApiHandler{
		redisClient: redis.NewClient(opt),
	}

	return handler, nil
}
