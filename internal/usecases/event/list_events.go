package event

import "github.com/AAteddy/event-booking-api/internal/domain/repositories"

type ListEventsUseCase struct{ Repo repositories.EventRepository }
