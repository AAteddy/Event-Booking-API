package persistance

import (
	"event-booking-api/internal/domain/entities"
	"event-booking-api/internal/domain/repositories"
	"gorm.io/gorm"
)

type TicketRepositoryImpl struct{ DB *gorm.DB }

func (repo *TicketRepositoryImpl) Save(ticket *entities.Ticket) error {
	return nil
}
