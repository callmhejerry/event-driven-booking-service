package bookingdto

type CreateBooking struct {
	RoomId    string `json:"room_id"`
	UserEmail string `json:"user_email"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
