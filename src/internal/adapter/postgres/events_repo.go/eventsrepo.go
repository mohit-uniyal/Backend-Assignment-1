package eventsrepo

import (
	"context"
	"event-booking/src/internal/core/model"
	outputport "event-booking/src/internal/port/output"
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
