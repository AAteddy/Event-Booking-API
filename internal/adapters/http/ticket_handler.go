package http

import (
	"net/http"

	"github.com/AAteddy/event-booking-api/internal/usecases/ticket"
	"github.com/go-playground/validator/v10"
)

type TicketHandler struct {
	BookUC   *ticket.BookTicketUseCase
	ListUC   *ticket.ListTicketsUseCase
	Validate *validator.Validate
}

func (h *TicketHandler) Book(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *TicketHandler) List(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
