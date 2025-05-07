package sessions

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

type UserSessionInfo struct {
	UserId      string `json:"userid"`
	AccessToken string `json:"accesstoken"`
}

var RedisContext = context.Background()

var RedisClient = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDISHOST"),
	Username: "default",
	Password: os.Getenv("REDISPASSWORD"),
	DB:       0,
})
