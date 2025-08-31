package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AAteddy/event-booking-api/internal/usecases/event"
	"github.com/go-playground/validator/v10"
)

type EventHandler struct {
	CreateUC *event.CreateEventUseCase
	ListUC   *event.ListEventsUseCase
	Validate *validator.Validate
}

type CreateEventInput struct {
	Title        string    `json:"title" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Date         time.Time `json:"date" validate:"required"`
	Location     string    `json:"location" validate:"required"`
	TotalTickets int       `json:"total_tickets" validate:"required,gte=1"`
}

type ListEventsInput struct {
	Offset int       `json:"offset" validate:"gte=0"`
	Limit  int       `json:"limit" validate:"gte=1,lte=100"`
	Date   time.Time `json:"date"`
}

func (h *EventHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateEventInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(string) // Set by JWT middleware
	event := event.CreateEventInput{
		Title:        input.Title,
		Description:  input.Description,
		Date:         input.Date,
		Location:     input.Location,
		TotalTickets: input.TotalTickets,
		OrganizerID:  userID,
	}

	createdEvent, err := h.CreateUC.Execute(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Event created successfully",
		"event":   createdEvent,
	})
}

func (h *EventHandler) List(w http.ResponseWriter, r *http.Request) {
	input := event.ListEventsInput{
		Offset: 0,
		Limit:  10,
	}

	events, err := h.ListUC.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
