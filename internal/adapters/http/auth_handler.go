package http

import (
	"encoding/json"
	"net/http"

	"github.com/AAteddy/event-booking-api/internal/domain/entities"
	"github.com/AAteddy/event-booking-api/internal/usecases/user"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	RegisterUC *user.RegisterUserUseCase
	LoginUC    *user.LoginUserUseCase
	Validate   *validator.Validate
}

type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=user organizer admin"`
	Name     string `json:"name" validate:"required,min=4"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string         `json:"token"`
	User  *entities.User `json:"user"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &user.RegisterUserInput{
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
		Name:     input.Name,
	}

	createdUser, err := h.RegisterUC.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the created user without the password
	createdUser.Password = ""
	// return the created user and status code 201 and success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user":    createdUser,
	})

}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loginResp, err := h.LoginUC.Execute(user.LoginUserInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: loginResp.Token, User: loginResp.User})
}
