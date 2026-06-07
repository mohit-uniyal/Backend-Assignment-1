package eventsusecase

import (
	"context"
	"event-booking/src/internal/core/constants"
	"event-booking/src/internal/core/dto"
	"event-booking/src/internal/core/model"
	inputport "event-booking/src/internal/port/input"
	outputport "event-booking/src/internal/port/output"
	"fmt"
	"log"
	"time"
)

type EventsUsecaseImpl struct {
	eventsRepo outputport.EventsRepo
}

func NewEventsUsecase(eventsRepo outputport.EventsRepo) inputport.EventsUsecase {
	return &EventsUsecaseImpl{
		eventsRepo: eventsRepo,
	}
}

func (e *EventsUsecaseImpl) CreateEvent(ctx context.Context, event *dto.Event) (*dto.Event, error) {
	// Validate the event
	if event.EventName == "" {
		return nil, fmt.Errorf("event name cannot be empty")
	}
	if event.TotalTickets == 0 {
		return nil, fmt.Errorf("total tickets cannot be 0")
	}

	eventTime, err := time.Parse(constants.ApplicationInputTimeFormat, event.EventTime)
	if err != nil {
		return nil, fmt.Errorf("invalid time layout: %s, expected: %s", event.EventTime, constants.ApplicationInputTimeFormat)
	}

	if eventTime.Before(time.Now().Add(24 * time.Hour)) {
		return nil, fmt.Errorf("event time should be at least 1 day ahead of the creation of event")
	}

	// Create the event
	eventId, err := e.eventsRepo.CreateEvent(ctx, &model.Event{
		EventName:    event.EventName,
		EventTime:    eventTime,
		TotalTickets: event.TotalTickets,
	})
	if err != nil {
		log.Printf("failed to create an event, %v", err)
		return nil, err
	}

	event.EventId = eventId
	return event, nil
}
