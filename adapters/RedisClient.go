package adapters

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"os"
)

var client *redis.Client

type UserDetails struct {
	user            string
	currentEstimate int
}

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

func NewRoom(ctx context.Context, user string) string {
	uid := uuid.New().String()

	//Store in a set a new room
	client.SAdd(ctx, "rooms", uid)

	//Add the current user to the newly created room in a hash
	client.HSet(ctx, uid, UserDetails{user: user, currentEstimate: 0})

	return uid
}
