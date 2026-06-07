package postgres

import (
	"context"
	"event-booking/src/internal/config"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetConnectionString(cfg *config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name,
	)
}

func NewPostgres(connectionString string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		log.Printf("failed to establish connection: %v", err)
		return nil, err
	}
	return db, nil
}
