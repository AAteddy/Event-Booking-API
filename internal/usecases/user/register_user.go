package user

import (
	"errors"

	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"github.com/AAteddy/event-booking-api/internal/domain/repositories"
	"github.com/go-playground/validator/v10"
)

type RegisterUserInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	Role     string `validate:"required,oneof=user organizer admin"`
	Name     string `validate:"required"`
}

type RegisterUserUseCase struct {
	Repo     repositories.UserRepository
	Validate *validator.Validate
}

func (uc *RegisterUserUseCase) Execute(input *RegisterUserInput) (*entities.User, error) {
	if err := uc.Validate.Struct(input); err != nil {
		return nil, err
	}

	// Check if email exists
	if _, err := uc.Repo.FindByEmail(input.Email); err == nil {
		return nil, errors.New("email already in use")
	}

	user := &entities.User{
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
		Name:     input.Name,
	}

	// Save user to repository
	if err := uc.Repo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}
