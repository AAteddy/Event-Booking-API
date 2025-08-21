package main

import (
	"event-booking-api/internal/adapters/http"
	"event-booking-api/internal/config"
	"event-booking-api/internal/infrastructure/cache"
	"event-booking-api/internal/infrastructure/persistence"
	"event-booking-api/internal/usecases/event"
	"event-booking-api/internal/usecases/ticket"
	"event-booking-api/internal/usecases/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize database
	db, err := persistence.SetupDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Redis
	redisClient, err := cache.NewRedisClient(cfg.RedisAddr, "", 0)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize validator
	validate := validator.New()

	// Repositories
	userRepo := &persistence.UserRepositoryImpl{DB: db}
	eventRepo := &persistence.EventRepositoryImpl{DB: db}
	ticketRepo := &persistence.TicketRepositoryImpl{DB: db}
	authRepo := &cache.AuthRepositoryImpl{Cache: redisClient}

	// Use cases
	registerUC := &user.RegisterUserUseCase{Repo: userRepo}
	loginUC := &user.LoginUserUseCase{Repo: userRepo}
	createEventUC := &event.CreateEventUseCase{Repo: eventRepo}
	listEventsUC := &event.ListEventsUseCase{Repo: eventRepo}
	bookTicketUC := &ticket.BookTicketUseCase{Repo: ticketRepo}
	listTicketsUC := &ticket.ListTicketsUseCase{Repo: ticketRepo}

	// Handlers
	authHandler := &http.AuthHandler{RegisterUC: registerUC, LoginUC: loginUC, Validate: validate}
	eventHandler := &http.EventHandler{CreateUC: createEventUC, ListUC: listEventsUC, Validate: validate}
	ticketHandler := &http.TicketHandler{BookUC: bookTicketUC, ListUC: listTicketsUC, Validate: validate}

	// Router
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	// Routes will be added in later sub-parts
	http.ListenAndServe(":8080", r)
}
