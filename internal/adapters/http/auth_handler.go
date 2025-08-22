package http


import (
	"github.com/AAteddy/event-booking-api/internal/usecases/user"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	RegisterUC *user.RegisterUserUseCase LoginUC *user.LoginUserUseCase Validate *validator.Validate
}