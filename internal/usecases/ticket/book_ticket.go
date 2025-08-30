package ticket

import (
	"errors"

	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"github.com/AAteddy/event-booking-api/internal/domain/repositories"
	"github.com/go-playground/validator/v10"
)

type BookTicketInput struct {
	EventID string `validate:"required"`
	UserID  string `validate:"required"`
}

type BookTicketUseCase struct {
	Repo      repositories.TicketRepository
	EventRepo repositories.EventRepository
	Validate  *validator.Validate
}

func (uc *BookTicketUseCase) Execute(input BookTicketInput) (*entities.Ticket, error) {
	if err := uc.Validate.Struct(input); err != nil {
		return nil, err
	}

	// Check event exists
	if _, err := uc.EventRepo.FindByID(input.EventID); err != nil {
		return nil, errors.New("event not found")
	}

	// Book ticket (handles concurrency)
	ticket, err := uc.Repo.BookTicket(input.EventID, input.UserID)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}
