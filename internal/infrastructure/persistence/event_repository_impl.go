package persistence

import (
	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct{ DB *gorm.DB }

func (repo *EventRepositoryImpl) Save(event *entities.Event) error {
	return repo.DB.Create(event).Error
}

func (repo *EventRepositoryImpl) FindByID(id string) (*entities.Event, error) {
	var event entities.Event
	err := repo.DB.First(&event, "id = ?", id).Error
	return &event, err
}

func (repo *EventRepositoryImpl) FindAll(offset, limit int, filters map[string]interface{}) ([]*entities.Event, error) {
	var events []*entities.Event
	query := repo.DB.Offset(offset).Limit(limit).Preload("Organizer")
	for key, value := range filters {
		if key == "date" {
			query = query.Where("date = ?", value)
		}
		// Add more filters as needed (e.g., category)

	}
	err := query.Find(&events).Error
	return events, err
}

func (repo *EventRepositoryImpl) Update(event *entities.Event) error {
	return repo.DB.Save(event).Error
}

func (repo *EventRepositoryImpl) Delete(id string) error {
	return repo.DB.Where("id = ?", id).Delete(&entities.Event{}).Error
}

func (repo *EventRepositoryImpl) Count(filters map[string]interface{}) (int64, error) {
	var count int64
	query := repo.DB.Model(&entities.Event{})
	for key, value := range filters {
		if key == "date" {
			query = query.Where("date = ?", value)
		}
	}
	err := query.Count(&count).Error
	return count, err
}
