package eventsrepo

import (
	"context"
	"database/sql"
	"event-booking/src/internal/core/model"
	outputport "event-booking/src/internal/port/output"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type EventsRepoImpl struct {
	db *pgxpool.Pool
}

func NewEventsRepo(db *pgxpool.Pool) outputport.EventsRepo {
	return &EventsRepoImpl{
		db: db,
	}
}

func (e *EventsRepoImpl) CreateEvent(ctx context.Context, event *model.Event) (int, error) {
	query := `INSERT INTO events(event_name, time, total_tickets) VALUES ($1, $2, $3) RETURNING event_id`

	row := e.db.QueryRow(ctx, query, event.EventName, event.EventTime, event.TotalTickets)

	var eventId int

	if err := row.Scan(&eventId); err != nil {
		log.Printf("failed to create event: %v", err)
		return 0, err
	}

	return eventId, nil

}

func (e *EventsRepoImpl) FetchAllFutureEvents(ctx context.Context) ([]*model.Event, error) {
	query := `
			SELECT
				event_id,
				event_name,
				time,
				total_tickets,
				tickets_sold
			FROM events
			WHERE time > NOW()
	`

	rows, err := e.db.Query(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no future events")
			fmt.Println("no events")
			return nil, nil
		}
		log.Printf("failed to fetch all future events: %v", err)
		return nil, err
	}

	var events []*model.Event

	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.EventId, &event.EventName, &event.EventTime, &event.TotalTickets, &event.TicketsSold); err != nil {
			log.Printf("failed to scan event: %v", err)
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}
