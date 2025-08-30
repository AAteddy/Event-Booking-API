package repositories

import "github.com/AAteddy/event-booking-api/internal/domain/entities"

type TicketRepository interface {
	Save(ticket *entities.Ticket) error
	FindByID(id string) (*entities.Ticket, error)
	FindByUserID(userID string, offset, limit int) ([]*entities.Ticket, error)
	Update(ticket *entities.Ticket) error
	Delete(id string) error
	CountByUserEventID(eventID string) (int64, error)
	CountByUserID(userID string) (int64, error)
	BookTicket(eventID, userID string) (*entities.Ticket, error)
}
