package ticket

import "github.com/AAteddy/event-booking-api/internal/domain/repositories"

type BookTicketUseCase struct{ Repo repositories.TicketRepository }
