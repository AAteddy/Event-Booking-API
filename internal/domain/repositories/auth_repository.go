package repositories

import "time"

type AuthRepository interface {
	BlacklistToken(token string, expiry time.Time) error
	IstokenBlacklisted(token string) (bool, error)
}
