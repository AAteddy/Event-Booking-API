package user

import "event-booking-api/internal/domain/repositories"

type RegisterUserUseCase struct{ Repo repositories.UserRepository }
