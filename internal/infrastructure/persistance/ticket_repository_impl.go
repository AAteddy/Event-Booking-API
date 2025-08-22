package persistance

import (
	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"github.com/AAteddy/event-booking-api/internal/domain/repositories"
	"gorm.io/gorm"
)

type TicketRepositoryImpl struct{ DB *gorm.DB }

func (repo *TicketRepositoryImpl) Save(ticket *entities.Ticket) error {
	return nil
}
