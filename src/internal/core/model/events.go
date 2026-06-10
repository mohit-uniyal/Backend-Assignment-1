package model

import "time"

type Event struct {
	EventId      int       `db:"event_id"`
	EventName    string    `db:"event_name"`
	EventTime    time.Time `db:"time"`
	TotalTickets int       `db:"total_tickets"`
	TicketsSold  int       `db:"tickets_sold"`
}
