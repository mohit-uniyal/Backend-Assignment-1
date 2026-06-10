package redisrepo

import (
	"context"
	outputport "event-booking/src/internal/port/output"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepoImpl struct {
	redisClient *redis.Client
}

func NewRedisRepo(redisClient *redis.Client) outputport.Cache {
	return &RedisRepoImpl{
		redisClient: redisClient,
	}
}

func (r *RedisRepoImpl) SetInt(ctx context.Context, key string, value int, expiration time.Duration) error {
	err := r.redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Printf("failed to set key %s, for value %d", key, value)
		return err
	}
	return nil
}
