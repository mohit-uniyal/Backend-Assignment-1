package main

import (
	"context"
	"event-booking/src/internal/adapter/postgres"
	eventsrepo "event-booking/src/internal/adapter/postgres/eventsrepo"
	"event-booking/src/internal/adapter/redis"
	"event-booking/src/internal/adapter/redis/redisrepo"
	"event-booking/src/internal/config"
	"event-booking/src/internal/input/handlers"
	"event-booking/src/internal/input/routes"
	cacheservice "event-booking/src/internal/usecase/cache"
	eventservice "event-booking/src/internal/usecase/events"
	"event-booking/src/migrations"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Load Configs
	cfg, err := config.LoadConfig(".secrets/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Load Migrations
	err = migrations.RunMigrations(postgres.GetConnectionString(cfg))
	if err != nil {
		log.Fatal(err)
	}

	// Intialize Databases
	db, err := postgres.NewPostgres(postgres.GetConnectionString(cfg))
	if err != nil {
		log.Fatal(err)
	}
	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Repo Functions
	cacheRepo := redisrepo.NewRedisRepo(redisClient)
	eventsRepo := eventsrepo.NewEventsRepo(db)

	// Initialize Usecase Functions
	cacheUsecase := cacheservice.NewCacheService(cacheRepo, eventsRepo)
	if err := cacheUsecase.PopulateEvents(context.Background()); err != nil {
		log.Printf("failed to populate events: %v", err)
	}
	eventsUsecase := eventservice.NewEventsUsecase(eventsRepo, cacheUsecase)

	// Initialize Handler Functions
	eventHandler := handlers.NewEventHandler(eventsUsecase)

	router := routes.NewRouter(eventHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	// Start server
	go func() {
		log.Printf("Server starting on port %d", cfg.Server.Port)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	// wait for shutdown signal
	sigChan := make(chan os.Signal, 1)

	signal.Notify(
		sigChan,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-sigChan

	log.Println("Shutdown signal received")

	// Give active requests time to finish
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Graceful shutdown fail: %v", err)

		if err := server.Close(); err != nil {
			log.Printf("failed to force close: %v", err)
		}
	}

	log.Println("Server stopped")

}
