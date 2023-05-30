package storage

import (
	"os"

	"github.com/gofiber/storage/redis/v2"
)

var RedisStore *redis.Storage

func ConnectToRedis() {
	store := redis.New(redis.Config{
		URL:   os.Getenv("REDIS_URL"),
		Reset: false,
	})
	RedisStore = store
}
