package eventservice

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
	eventsRepo   outputport.EventsRepo
	cacheService inputport.CacheUsecase
}

func NewEventsUsecase(eventsRepo outputport.EventsRepo, cacheService inputport.CacheUsecase) inputport.EventsUsecase {
	return &EventsUsecaseImpl{
		eventsRepo:   eventsRepo,
		cacheService: cacheService,
	}
}

func (e *EventsUsecaseImpl) CreateEvent(ctx context.Context, event *dto.Event) (*dto.Event, error) {
	// Validate the event
	if event == nil {
		log.Printf("event is nil in CreateEvent")
		return nil, fmt.Errorf("event is nil in CreateEvent")
	}
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

	//push the event to cache
	err = e.cacheService.SetEvent(ctx, &model.Event{
		EventId:      eventId,
		EventName:    event.EventName,
		EventTime:    eventTime,
		TicketsSold:  event.TicketsSold,
		TotalTickets: event.TotalTickets,
	})
	if err != nil {
		log.Printf("failed to set event in cache: %v", err)
	}

	event.EventId = eventId
	return event, nil
}
