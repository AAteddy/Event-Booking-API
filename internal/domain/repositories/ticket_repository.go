package repositories

import "event-booking-api/internal/domain/entities"

type TicketRepository interface {
	Save(ticket *entities.Ticket) error
}
