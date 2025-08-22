package cache

import (
	"github.com/AAteddy/event-booking-api/internal/domain/repositories"
	"time"
)

type AuthRepositoryImpl struct{ Cache *RedisClient }

func (repo *AuthRepositoryImpl) BlacklistToken(token string, expiry time.Time) error {
	return nil
}
