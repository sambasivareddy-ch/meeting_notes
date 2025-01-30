package sessions

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type UserSessionInfo struct {
	UserId      string `json:"userid"`
	AccessToken string `json:"accesstoken"`
}

var RedisContext = context.Background()

var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})
