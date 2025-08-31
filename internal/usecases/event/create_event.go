package event

import (
	"errors"
	"time"

	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"github.com/AAteddy/event-booking-api/internal/domain/repositories"
	"github.com/go-playground/validator/v10"
)

type CreateEventInput struct {
	Title        string `validate:"required"`
	Description  string
	Date         time.Time `validate:"required"`
	TotalTickets int       `validate:"required,gte=0"`
	Location     string    `validate:"required"`
	OrganizerID  string    `validate:"required"`
}

type CreateEventUseCase struct {
	Repo     repositories.EventRepository
	UserRepo repositories.UserRepository
	Validate *validator.Validate
}

func (uc *CreateEventUseCase) Execute(input CreateEventInput) (*entities.Event, error) {
	if err := uc.Validate.Struct(input); err != nil {
		return nil, err
	}

	// Check if user with Role organizer exists and has role "organizer" or admin
	user, err := uc.UserRepo.FindByID(input.OrganizerID)
	if err != nil || (user.Role != "organizer" && user.Role != "admin") {
		return nil, errors.New("only organizers or admins can create events")
	}

	event := &entities.Event{
		Title:            input.Title,
		Description:      input.Description,
		Date:             input.Date,
		TotalTickets:     input.TotalTickets,
		AvailableTickets: input.TotalTickets,
		OrganizerID:      input.OrganizerID,
	}

	if err := uc.Repo.Save(event); err != nil {
		return nil, err
	}

	return event, nil
}
