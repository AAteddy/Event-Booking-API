package event

import "event-booking-api/internal/domain/repositories"

type CreateEventUseCase struct{ Repo repositories.EventRepository }
