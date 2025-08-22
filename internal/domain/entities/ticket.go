package entities

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct{
	ID        		string         		`gorm:"type:uuid;primaryKey"`
    EventID   		string         		`gorm:"type:uuid;not null"`
    Event     		Event          		`gorm:"foreignKey:EventID"`
    UserID    		string         		`gorm:"type:uuid;not null"`
    User      		User           		`gorm:"foreignKey:UserID"`
    Status    		string         		`gorm:"not null;check:status IN ('booked','canceled')"`
    CreatedAt 		time.Time
    UpdatedAt 		time.Time
    DeletedAt 		gorm.DeletedAt 		`gorm:"index"`
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()
	t.Status = "booked"
	return nil
}
