package booking

import (
	"context"
	"errors"
	"time"

	"github.com/callmhejerry/event-driven-booking-service/internal/db"
	bookingdto "github.com/callmhejerry/event-driven-booking-service/internal/dto/booking_dto"
)

type BookingStatus string

var (
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
)

type BookingService struct {
	database *db.Database
}

func NewBookingService(database *db.Database) *BookingService {
	return &BookingService{
		database: database,
	}
}

func (s *BookingService) CreateBooking(
	ctx context.Context, request bookingdto.CreateBooking,
) (*db.Booking, error) {
	startTime, err := time.Parse(time.RFC3339, request.StartTime)
	if err != nil {
		return nil, errors.New("invalid start_time")
	}

	endTime, err := time.Parse(time.RFC3339, request.EndTime)
	if err != nil {
		return nil, errors.New("invalid end_time")
	}

	if startTime.After(endTime) {
		return nil, errors.New("start_time must come before end_time")
	}

	booking, err := s.database.CreateBooking(ctx, db.CreateBookingParams{
		RoomID:    request.RoomId,
		StartTime: startTime,
		EndTime:   endTime,
		UserEmail: request.UserEmail,
		Status:    string(BookingStatusConfirmed),
	})

	if err != nil {
		return nil, err
	}

	return &booking, nil
}
