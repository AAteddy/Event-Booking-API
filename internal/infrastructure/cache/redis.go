package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string, db int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Test connection
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &RedisClient{Client: client}, nil
}

func (c *RedisClient) Close() error {
	return c.Client.Close()
}
