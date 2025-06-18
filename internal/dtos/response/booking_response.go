package response

import "time"

type BookingResponse struct {
	ID           uint      `json:"id"`
	GuestID      uint      `json:"guest_id"`
	RoomID       uint      `json:"room_id"`
	CheckinDate  time.Time `json:"checkin_date"`
	CheckoutDate time.Time `json:"checkout_date"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
