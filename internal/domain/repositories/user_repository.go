package repositories

import "event-booking-api/internal/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
}
