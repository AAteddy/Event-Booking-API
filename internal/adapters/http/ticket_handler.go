package http

import (
	"encoding/json"
	"net/http"

	"github.com/AAteddy/event-booking-api/internal/usecases/ticket"
	"github.com/go-playground/validator/v10"
)

type TicketHandler struct {
	BookUC   *ticket.BookTicketUseCase
	ListUC   *ticket.ListTicketsUseCase
	Validate *validator.Validate
}

type BookTicketInput struct {
	EventID string `json:"event_id" validate:"required,uuid"`
}

type ListTicketsInput struct {
	Offset int `json:"offset" validate:"gte=0"`
	Limit  int `json:"limit" validate:"gte=1,lte=100"`
}

func (h *TicketHandler) Book(w http.ResponseWriter, r *http.Request) {
	var input BookTicketInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(string)
	ticket := ticket.BookTicketInput{
		EventID: input.EventID,
		UserID:  userID,
	}

	bookedTicket, err := h.BookUC.Execute(ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Ticket booked successfully",
		"ticket":  bookedTicket,
	})
}

func (h *TicketHandler) List(w http.ResponseWriter, r *http.Request) {
	var input ListTicketsInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(string)
	response, err := h.ListUC.Execute(ticket.ListTicketsInput{
		UserID: userID,
		Offset: input.Offset,
		Limit:  input.Limit,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tickets": response.Tickets,
		"total":   response.Total,
	})
}
