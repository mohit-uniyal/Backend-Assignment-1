package routes

import (
	"event-booking/src/internal/input/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter(
	eventHandler *handlers.EventHandler,
) *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/", eventHandler.HealthCheck)

	//Events Routes
	eventsRouter := apiRouter
	eventsRouter.Post("/", eventHandler.CreateEventHandler)
	eventsRouter.Mount("/events", apiRouter)

	// Add prefix
	apiRouter.Mount("/api", apiRouter)

	return apiRouter
}
