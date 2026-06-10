package redis

import (
	"context"
	"event-booking/src/internal/config"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Printf("failed to connect to redis: %v", err)
		return nil, err
	}

	return redisClient, nil
}
