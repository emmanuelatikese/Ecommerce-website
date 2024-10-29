package redis_db

import (
	"context"
	"os"
	"github.com/go-redis/redis/v8"
	"github.com/subosito/gotenv"
)

var RedisCli *redis.Client

func NewClient() error {
	ctx := context.TODO()
	gotenv.Load()
	url := os.Getenv("REDIS_URL")

	RedisCli = redis.NewClient(&redis.Options{
		Addr:url,
	})
	_, err := RedisCli.Ping(ctx).Result()
	if err != nil{
		return err
	}
	return nil
}