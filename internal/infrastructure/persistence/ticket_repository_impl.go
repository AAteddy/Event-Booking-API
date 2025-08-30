package persistence

import (
	"errors"

	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"gorm.io/gorm"
)

type TicketRepositoryImpl struct{ DB *gorm.DB }

func (repo *TicketRepositoryImpl) Save(ticket *entities.Ticket) error {
	return repo.DB.Create(ticket).Error
}

func (repo *TicketRepositoryImpl) FindByID(id string) (*entities.Ticket, error) {
	var ticket entities.Ticket
	err := repo.DB.Preload("Event").Preload("User").First(&ticket, "id = ?", id).Error
	return &ticket, err
}

func (repo *TicketRepositoryImpl) FindByUserID(userID string, offset, limit int) ([]*entities.Ticket, error) {
	var tickets []*entities.Ticket
	err := repo.DB.Preload("Event").Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&tickets).Error
	return tickets, err
}

func (repo *TicketRepositoryImpl) Update(ticket *entities.Ticket) error {
	return repo.DB.Save(ticket).Error
}

func (repo *TicketRepositoryImpl) Delete(id string) error {
	return repo.DB.Where("id = ?", id).Delete(&entities.Ticket{}).Error
}

func (repo *TicketRepositoryImpl) CountByUserEventID(eventID string) (int64, error) {
	var count int64
	err := repo.DB.Model(&entities.Ticket{}).Where("event_id = ? AND status = ?", eventID, "booked").Count(&count).Error
	return count, err
}

func (repo *TicketRepositoryImpl) CountByUserID(userID string) (int64, error) {
	var count int64
	err := repo.DB.Model(&entities.Ticket{}).Where("user_id = ? AND status = ?", userID, "booked").Count(&count).Error
	return count, err
}

func (repo *TicketRepositoryImpl) BookTicket(eventID, userID string) (*entities.Ticket, error) {
	return nil, repo.DB.Transaction(func(tx *gorm.DB) error {
		var event entities.Event
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&event, "id = ?", eventID).Error; err != nil {
			return err
		}
		if event.AvailableTickets < 1 {
			return errors.New("no tickets available")
		}
		ticket := &entities.Ticket{
			EventID: eventID,
			UserID:  userID,
			Status:  "booked",
		}
		if err := tx.Create(ticket).Error; err != nil {
			return err
		}
		event.AvailableTickets--
		if err := tx.Save(&event).Error; err != nil {
			return err
		}
		return nil
	})
}
