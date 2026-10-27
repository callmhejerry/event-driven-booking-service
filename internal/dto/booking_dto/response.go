package bookingdto

type BookingResponseDto struct {
	ID        string `json:"id"`
	RoomId    string `json:"room_id"`
	UserEmail string `json:"user_email"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
