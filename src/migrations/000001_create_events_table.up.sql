CREATE TABLE IF NOT EXISTS events(
    event_id SERIAL PRIMARY KEY,
    event_name VARCHAR(255) NOT NULL,
    time TIME NOT NULL,
    total_tickets INTEGER DEFAULT 0,
    tickets_sold INTEGER DEFAULT 0
);