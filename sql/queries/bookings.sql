-- name: GetAllBookings :many
SELECT * FROM bookings;


-- name: CreateBooking :one
INSERT INTO bookings (
    room_id , 
    start_time, 
    end_time,
    user_email,
    status
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


