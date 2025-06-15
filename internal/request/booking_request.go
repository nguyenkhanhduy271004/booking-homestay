package request

import "time"

type BookingRequest struct {
	CheckInDate  time.Time `json:"check_in_date" binding:"required"`
	CheckOutDate time.Time `json:"check_out_date" binding:"required"`
	GuestID      uint      `json:"guest_id" binding:"required"`
	RoomID       uint      `json:"room_id" binding:"required"`
	PaymentID    uint      `json:"payment_id" binding:"required"`
	Status       string    `json:"status" binding:"required"`
}
