package http

import (
	"net/http"

	"github.com/AAteddy/event-booking-api/internal/usecases/event"
	"github.com/go-playground/validator/v10"
)

type EventHandler struct {
	CreateUC *event.CreateEventUseCase
	ListUC   *event.ListEventsUseCase
	Validate *validator.Validate
}

func (h *EventHandler) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *EventHandler) List(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
