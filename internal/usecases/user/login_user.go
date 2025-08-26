package user

import (
	"errors"
	"time"

	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"github.com/AAteddy/event-booking-api/internal/domain/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)


type LoginUserInput struct {
	Email string `validate:"required,email"`
	Password string `validate:"required"`
}

type LoginUserUseCase struct{ 
	Repo repositories.UserRepository
	Validate *validator.Validate
	JWTSecret string
}

type LoginResponse struct {
	Token string
	User *entities.User
}

func (uc *LoginUserUseCase) Execute(input LoginUserInput) (*LoginResponse, error) {
	if err := uc.Validate.Struct(input); err != nil {
		return nil, err
	}

	user, err := uc.Repo.FindByEmail(input.Email)
	if err != nil || !user.CheckPassword(input.Password) {
		return nil, errors.New("invalid credentials")
	}
	
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(uc.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: tokenString,
		User:  user,
	}, nil
}