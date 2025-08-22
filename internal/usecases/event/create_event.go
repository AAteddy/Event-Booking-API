package event

import "github.com/AAteddy/event-booking-api/internal/domain/repositories"

type CreateEventUseCase struct{ Repo repositories.EventRepository }
