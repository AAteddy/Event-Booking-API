package cache

import "github.com/redis/go-redis/v9"

type RedisClient struct{ Client *redis.Client }

func NewRedisClient(addr, password string, db int) (*RedisClient, error) {
	return nil, nil
}
