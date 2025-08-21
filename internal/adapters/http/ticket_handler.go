package http

import (
	"event-booking-api/internal/usecases/ticket"
	"github.com/go-playground/validator/v10"
)

type TicketHandler struct {
	BookUC *ticket.BookTicketUseCase ListUC *ticket.ListTicketsUseCase Validate *validator.Validate
}
