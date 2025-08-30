package persistence

import (
	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto-migrate schemas
	err = db.AutoMigrate(&entities.User{}, &entities.Event{}, &entities.Ticket{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
