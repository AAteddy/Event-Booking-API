package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthRepositoryImpl struct{ 
	Cache *RedisClient 
}

func (repo *AuthRepositoryImpl) BlacklistToken(token string, expiry time.Time) error {
	ctx := context.Background()
	return repo.Cache.Client.Set(ctx, "blacklist:"+token, "1", time.Until(expiry)).Err()
}

func (repo *AuthRepositoryImpl) IsTokenBlacklisted(token string) (bool, error) {
	ctx := context.Background()
	_, err := repo.Cache.Client.Get(ctx, "blacklist:"+token).Result()
	if err == redis.Nil {
		return false, nil
	}
	return err == nil, err
}
