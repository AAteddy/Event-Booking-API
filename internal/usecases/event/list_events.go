package event

import "event-booking-api/internal/domain/repositories"

type ListEventsUseCase struct{ Repo repositories.EventRepository }
