package ticket

import "event-booking-api/internal/domain/repositories"

type BookTicketUseCase struct{ Repo repositories.TicketRepository }
