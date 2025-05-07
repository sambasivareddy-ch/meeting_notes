package sessions

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type UserSessionInfo struct {
	UserId      string `json:"userid"`
	AccessToken string `json:"accesstoken"`
}

var (
	RedisContext = context.Background()
	RedisClient  *redis.Client
)

func init() {
	redisURL := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	RedisClient = redis.NewClient(opt)
}
