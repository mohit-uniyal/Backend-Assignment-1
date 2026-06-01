package routes

import (
	"event-booking/src/internal/input/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter(
	eventHandler *handlers.EventHandler,
) *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/", handlers.HealthCheck)

	// Add prefix
	apiRouter.Mount("/api", apiRouter)

	return apiRouter
}
