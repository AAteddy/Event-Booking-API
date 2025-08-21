package persistance

import (
	"event-booking-api/internal/domain/entities"
	"event-booking-api/internal/domain/repositories"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct{ DB *gorm.DB }

func (repo *EventRepositoryImpl) Save(event *entities.Event) error {
	return nil
}
