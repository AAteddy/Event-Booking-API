package http

import (
	"net/http"

	"github.com/AAteddy/event-booking-api/internal/usecases/user"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	RegisterUC *user.RegisterUserUseCase
	LoginUC    *user.LoginUserUseCase
	Validate   *validator.Validate
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
