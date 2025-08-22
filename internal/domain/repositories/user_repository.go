package repositories

import "github.com/AAteddy/event-booking-api/internal/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
	FindByID(id string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindAll(offset, limit int) ([]*entities.User, error)
	Update(user *entities.User) error
	Delete(id string) error
	Count() (int64, error)
}
