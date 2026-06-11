package cacheservice

import (
	"context"
	"event-booking/src/internal/core/model"
	inputport "event-booking/src/internal/port/input"
	outputport "event-booking/src/internal/port/output"
	"fmt"
	"log"
	"time"
)

type CacheUsecaseImpl struct {
	cacheRepo  outputport.Cache
	eventsRepo outputport.EventsRepo
}

func NewCacheService(cacheRepo outputport.Cache, eventsRepo outputport.EventsRepo) inputport.CacheUsecase {
	return &CacheUsecaseImpl{
		cacheRepo:  cacheRepo,
		eventsRepo: eventsRepo,
	}
}

func (c *CacheUsecaseImpl) PopulateEvents(ctx context.Context) error {

	events, err := c.eventsRepo.FetchAllFutureEvents(ctx)
	if err != nil {
		log.Printf("failed to fetch future events: %v", err)
		return err
	}

	for _, event := range events {
		eventKey := getEventTicketsKey(event)
		if eventKey == "" {
			log.Printf("failed to form a key for the event: %d", event.EventId)
			continue
		}

		err := c.cacheRepo.SetInt(ctx, eventKey, event.TotalTickets-event.TicketsSold, time.Until(event.EventTime))
		if err != nil {
			log.Printf("failed to set event in cache, %v", err)
		}

	}

	return nil
}

func (c *CacheUsecaseImpl) SetEvent(ctx context.Context, event *model.Event) error {
	if event == nil {
		log.Printf("event is nil")
		return fmt.Errorf("event is nil in SetEvent")
	}
	eventKey := getEventTicketsKey(event)
	if eventKey == "" {
		log.Printf("failed to form a key for the event: %d", event.EventId)
		return fmt.Errorf("failed to form a key for the event: %d", event.EventId)
	}

	err := c.cacheRepo.SetInt(ctx, eventKey, event.TotalTickets-event.TicketsSold, time.Until(event.EventTime))
	if err != nil {
		log.Printf("failed to set event in cache, %v", err)
		return err
	}

	return nil
}
