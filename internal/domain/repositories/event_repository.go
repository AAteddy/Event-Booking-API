package repositories

import "event-booking-api/internal/domain/entities"

type EventRepository interface {
	Save(event *entities.Event) error
}
