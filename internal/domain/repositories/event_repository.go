package repositories

import (
	"github.com/AAteddy/event-booking-api/internal/domain/entities"
)

type EventRepository interface {
	Save(event *entities.Event) error
	FindByID(id string) (*entities.Event, error)
	FindAll(offset, limit int, fileters map[string]interface{}) ([]*entities.Event, error)
	Update(event *entities.Event) error
	Delete(id string) error 
	Count(filters map[string]interface{}) (int64, error)
}
