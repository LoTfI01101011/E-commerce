package internal

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Redis *redis.Client

func init() {
	LoadEnv()
}

func RedisConnection() {
	Redis = redis.NewClient(&redis.Options{
		Addr: "localhost:" + os.Getenv("RedisPort"),
	})
}
