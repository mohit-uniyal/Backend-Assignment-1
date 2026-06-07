package dto

type Event struct {
	EventId      int    `json:"event_id"`
	EventName    string `json:"event_name"`
	EventTime    string `json:"event_time"`
	TotalTickets int    `json:"total_tickets"`
	TicketsSold  int    `json:"tickets_sold"`
}
