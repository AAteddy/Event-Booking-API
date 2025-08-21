package user

import "event-booking-api/internal/domain/repositories"

type LoginUserUseCase struct{ Repo repositories.UserRepository }
