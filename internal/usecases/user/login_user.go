package user

import "github.com/AAteddy/event-booking-api/internal/domain/repositories"

type LoginUserUseCase struct{ Repo repositories.UserRepository }
