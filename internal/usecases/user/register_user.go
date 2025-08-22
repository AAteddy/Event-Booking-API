package user

import "github.com/AAteddy/event-booking-api/internal/domain/repositories"

type RegisterUserUseCase struct{ Repo repositories.UserRepository }
