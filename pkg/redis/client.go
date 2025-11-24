package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // si no hay password
		DB:       0,
	})
}
