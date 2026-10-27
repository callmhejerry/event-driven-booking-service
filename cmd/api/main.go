package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/callmhejerry/event-driven-booking-service/internal/config"
	"github.com/callmhejerry/event-driven-booking-service/internal/db"
	"github.com/gin-gonic/gin"
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

	_, err = db.NewDatabase(*dbConfig)
	if err != nil {
		log.Fatalf("Failed to create database %v\n", err)
	}

	router := gin.Default()

	httpServer := http.Server{
		Addr:              ":2000",
		Handler:           router,
		ReadTimeout:       time.Second * 15,
		WriteTimeout:      time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Sever is shutting down")
		} else {
			log.Fatalf("Failed to start server: %s", err)
		}
	}

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGTERM, syscall.SIGINT)
	defer close(shutdownChan)

	<-shutdownChan
	shutdownCtx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Failed to shutdown server")
	}
}
