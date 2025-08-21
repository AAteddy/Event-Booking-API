package persistance

import (
	"event-booking-api/internal/domain/entities"
	"event-booking-api/internal/domain/repositories"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{ DB *gorm.DB }

func (repo *UserRepositoryImpl) Save(user *entities.User) error {
	return nil
}
