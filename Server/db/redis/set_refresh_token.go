package redis_db

import (
	"context"
	"time"
)

func SetRfToken(key string, token interface{}, ctx context.Context) {
	RedisCli.Set(ctx, key, token, time.Hour * 24 * 7)
}