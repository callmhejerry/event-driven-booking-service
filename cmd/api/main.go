package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/callmhejerry/event-driven-booking-service/internal/application"
	"github.com/callmhejerry/event-driven-booking-service/internal/config"
	"github.com/callmhejerry/event-driven-booking-service/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env files")
	}
	dbConfig, err := config.NewDatabaseConfig()
	if err != nil {
		log.Fatal("Failed to create database configuration")
	}

	database, err := db.NewDatabase(*dbConfig)
	if err != nil {
		log.Fatalf("Failed to create database %v\n", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	application := application.NewApplication(dbConfig, logger, database)

	if err := application.Start(); err != nil {
		logger.Error("Failed to start server", "error", err.Error())
		os.Exit(1)
	}

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGTERM, syscall.SIGINT)
	defer close(shutdownChan)

	<-shutdownChan
	shutdownCtx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()

	if err := application.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Failed to shutdown server")
	}
}
