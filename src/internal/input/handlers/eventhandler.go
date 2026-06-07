package handlers

import (
	"encoding/json"
	"event-booking/src/internal/core/dto"
	inputport "event-booking/src/internal/port/input"
	"log"
	"net/http"
)

type EventHandler struct {
	eventsUsecase inputport.EventsUsecase
}

func NewEventHandler(eventsUsecase inputport.EventsUsecase) *EventHandler {
	return &EventHandler{
		eventsUsecase: eventsUsecase,
	}
}

func (e *EventHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is working"))
}

func (e *EventHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	// parse incoming request
	var req *dto.Event
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("faild to unmarshal request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// handle the request
	event, err := e.eventsUsecase.CreateEvent(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(event); err != nil {
		log.Printf("failed to encode the event: %v", err)
	}

}
