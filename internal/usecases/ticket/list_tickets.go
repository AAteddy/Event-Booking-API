package ticket

import (
	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"github.com/AAteddy/event-booking-api/internal/domain/repositories"
)

type ListTicketsInput struct {
	UserID string
	Offset int
	Limit  int
}

type ListTicketsResponse struct {
	Tickets []*entities.Ticket
	Total   int64
}

type ListTicketsUseCase struct {
	Repo repositories.TicketRepository
}

func (uc *ListTicketsUseCase) Execute(input ListTicketsInput) (*ListTicketsResponse, error) {
	tickets, err := uc.Repo.FindByUserID(input.UserID, input.Offset, input.Limit)
	if err != nil {
		return nil, err
	}

	total, err := uc.Repo.CountByUserID(input.UserID)
	if err != nil {
		return nil, err
	}

	return &ListTicketsResponse{
		Tickets: tickets,
		Total:   total,
	}, nil

}
