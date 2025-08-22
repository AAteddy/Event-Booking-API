package entities

import ( 
	"time" 
	"github.com/google/uuid" 
	"gorm.io/gorm" 
)


type Event struct{
	ID 					string 						`gorm:"type:uuid;primaryKey"`
	OrganizerID 		string 						`gorm:"type:uuid;not null"`
	Organizer 			User 						`gorm:"foreignKey:OrganizerID"`
	Title 				string 						`gorm:"not null"`
	Description 		string 
	Date 				time.Time 					`gorm:"not null"`
	TotalTickets 		int 						`gorm:"not null;check:total_tickets >= 0"`
	AvailableTickets 	int 						`gorm:"not null;check:available_tickets >= 0"`
	CreatedAt 			time.Time 
	UpdatedAt 			time.Time 
	DeletedAt 			gorm.DeletedAt 				`gorm:"index"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.New().String()
	e.AvailableTickets = e.TotalTickets
	return nil
}
