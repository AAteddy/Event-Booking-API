package http


import (
	"github.com/AAteddy/event-booking-api/internal/usecases/event"
	"github.com/go-playground/validator/v10"
)

type EventHandler struct {
	CreateUC *event.CreateEventUseCase ListUC *event.ListEventsUseCase Validate *validator.Validate
}
