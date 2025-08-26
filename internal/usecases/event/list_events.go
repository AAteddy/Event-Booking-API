package event

import (
	"time"

	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"github.com/AAteddy/event-booking-api/internal/domain/repositories"
)


type ListEventsInput struct {
	Offset int
	Limit  int
	Date   *time.Time
}

type ListEventsResponse struct {
	Events []*entities.Event
	Total  int64
}

type ListEventsUseCase struct{ 
	Repo repositories.EventRepository 
}

func (uc *ListEventsUseCase) Execute(input ListEventsInput) (*ListEventsResponse, error) {
	filters := make(map[string]interface{})
	if input.Date != nil {
		filters["date"] = *input.Date
	}

	events, err := uc.Repo.FindAll(input.Offset, input.Limit, filters)
	if err != nil {
		return nil, err
	}

	total, err := uc.Repo.Count(filters)
	if err != nil {
		return nil, err
	}
	
	return &ListEventsResponse{
		Events: events,
		Total:  total,
	}, nil
}