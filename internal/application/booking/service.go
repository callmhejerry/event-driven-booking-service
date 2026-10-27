package booking

import "github.com/callmhejerry/event-driven-booking-service/internal/db"

type BookingService struct {
	database *db.Database
}

func NewBookingService(database *db.Database) *BookingService {
	return &BookingService{
		database: database,
	}
}
