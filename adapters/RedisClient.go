package adapters

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var client *redis.Client

func CreateClient() (*redis.Client, error) {
	opt, err := redis.ParseURL(os.Getenv("REDIS_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}

	client = redis.NewClient(opt)
	return client, nil
}

func IncAndGetRefreshed(ctx context.Context) (int64, error) {
	return client.Incr(ctx, "refreshes").Result()
}
