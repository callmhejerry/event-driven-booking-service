package application

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/callmhejerry/event-driven-booking-service/internal/application/booking"
	"github.com/callmhejerry/event-driven-booking-service/internal/config"
	"github.com/callmhejerry/event-driven-booking-service/internal/db"
	"github.com/gin-gonic/gin"
)

type Application struct {
	DatabaseConfig *config.DatabaseConfig
	Logger         *slog.Logger
	httpServer     *http.Server

	BookingService *booking.BookingService
}

func NewApplication(
	databaseConfig *config.DatabaseConfig,
	logger *slog.Logger,
	database *db.Database,
) *Application {
	bookingService := booking.NewBookingService(database)

	app := &Application{
		DatabaseConfig: databaseConfig,
		Logger:         logger,
		BookingService: bookingService,
	}

	app.setup()

	return app
}

func (a *Application) setup() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	booking.NewBookingController(router, a.BookingService)

	httpServer := http.Server{
		Addr:              ":2000",
		Handler:           router,
		ReadTimeout:       time.Second * 15,
		WriteTimeout:      time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
	}

	a.httpServer = &httpServer
	return router
}

func (a *Application) Start() error {
	return a.httpServer.ListenAndServe()
}

func (a *Application) Shutdown(ctx context.Context) error {
	return a.httpServer.Shutdown(ctx)
}
