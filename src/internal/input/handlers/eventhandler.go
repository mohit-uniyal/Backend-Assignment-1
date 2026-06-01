package handlers

import "net/http"

type EventHandler struct {
}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is working"))
}
