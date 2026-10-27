package booking

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/callmhejerry/event-driven-booking-service/internal/config"
	"github.com/callmhejerry/event-driven-booking-service/internal/db"
	bookingdto "github.com/callmhejerry/event-driven-booking-service/internal/dto/booking_dto"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testcase struct {
	request        bookingdto.CreateBooking
	expectedResult db.Booking
	expectError    bool
}

func TestCreateBooking(t *testing.T) {
	err := godotenv.Load("../../../.env")

	require.NoError(t, err)
	config, err := config.NewDatabaseConfig()

	fmt.Printf("Config %+v\n", config)
	require.NoError(t, err)
	database, err := db.NewDatabase(*config)
	require.NoError(t, err)

	testcases := []testcase{
		{
			request: bookingdto.CreateBooking{
				RoomId:    "room-id",
				UserEmail: "user_email",
				StartTime: time.Now().Format(time.RFC3339),
				EndTime:   time.Now().Add(time.Minute * 5).Format(time.RFC3339),
			},
			expectedResult: db.Booking{
				RoomID:    "room-id",
				ID:        uuid.New(),
				UserEmail: "user_email",
				Status:    "confirmed",
			},
			expectError: false,
		},
		{
			request: bookingdto.CreateBooking{
				RoomId:    "room-id",
				UserEmail: "user_email",
				StartTime: time.Now().Add(time.Minute * 10).Format(time.RFC3339),
				EndTime:   time.Now().Format(time.RFC3339),
			},
			expectedResult: db.Booking{
				RoomID:    "room-id",
				ID:        uuid.New(),
				UserEmail: "user_email",
				Status:    "confirmed",
			},
			expectError: true,
		},
	}

	service := NewBookingService(database)

	for _, tc := range testcases {
		actualResult, err := service.CreateBooking(context.Background(), tc.request)

		if tc.expectError {
			assert.Error(t, err)
			continue
		}

		fmt.Printf("Actual Result %+v\n", actualResult)
		require.NoError(t, err)
		assert.Equal(t, tc.expectedResult.RoomID, actualResult.RoomID)
		assert.Equal(t, tc.expectedResult.Status, actualResult.Status)
	}
}
