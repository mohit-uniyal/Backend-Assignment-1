package main

import (
	"context"
	"event-booking/src/internal/config"
	"event-booking/src/internal/input/handlers"
	"event-booking/src/internal/input/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := config.LoadConfig(".secrets/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	eventHandler := handlers.NewEventHandler()

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
