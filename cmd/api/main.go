package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	httpHandlers "github.com/AAteddy/event-booking-api/internal/adapters/http"
	"github.com/AAteddy/event-booking-api/internal/config"
	"github.com/AAteddy/event-booking-api/internal/infrastructure/cache"
	"github.com/AAteddy/event-booking-api/internal/infrastructure/persistence"
	"github.com/AAteddy/event-booking-api/internal/usecases/event"
	"github.com/AAteddy/event-booking-api/internal/usecases/ticket"
	"github.com/AAteddy/event-booking-api/internal/usecases/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
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
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0
	}

	redisClient, err := cache.NewRedisClient(cfg.RedisAddr, os.Getenv("REDIS_PASSWORD"), redisDB)
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()

	// Initialize validator
	validate := validator.New()

	// Repositories
	userRepo := &persistence.UserRepositoryImpl{DB: db}
	eventRepo := &persistence.EventRepositoryImpl{DB: db}
	ticketRepo := &persistence.TicketRepositoryImpl{DB: db}
	authRepo := &cache.AuthRepositoryImpl{Cache: redisClient}

	// Use cases
	registerUC := &user.RegisterUserUseCase{Repo: userRepo, Validate: validate}
	loginUC := &user.LoginUserUseCase{Repo: userRepo, Validate: validate, JWTSecret: cfg.JWTSecret}
	createEventUC := &event.CreateEventUseCase{Repo: eventRepo, UserRepo: userRepo, Validate: validate}
	listEventsUC := &event.ListEventsUseCase{Repo: eventRepo}
	bookTicketUC := &ticket.BookTicketUseCase{Repo: ticketRepo, EventRepo: eventRepo, Validate: validate}
	listTicketsUC := &ticket.ListTicketsUseCase{Repo: ticketRepo}

	// Handlers
	authHandler := &httpHandlers.AuthHandler{RegisterUC: registerUC, LoginUC: loginUC, Validate: validate}
	eventHandler := &httpHandlers.EventHandler{CreateUC: createEventUC, ListUC: listEventsUC, Validate: validate}
	ticketHandler := &httpHandlers.TicketHandler{BookUC: bookTicketUC, ListUC: listTicketsUC, Validate: validate}

	// Router
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Temporary routes to use handlers and authRepo (to be expanded in Part 15.4)
	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)
	r.Post("/events", eventHandler.Create)
	r.Get("/events", eventHandler.List)
	r.Post("/tickets", ticketHandler.Book)
	r.Get("/tickets", ticketHandler.List)
	// Use authRepo to ensure it's not unused (temporary, will be used in middleware)
	if authRepo != nil {
		log.Println("Auth repository initialized")
	}

	// Routes will be added in later sub-parts
	http.ListenAndServe(":8080", r)
}
