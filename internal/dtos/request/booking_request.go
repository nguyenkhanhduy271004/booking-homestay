package request

import "time"

type BookingRequest struct {
	CheckinDate  time.Time `json:"checkin_date" binding:"required"`
	CheckoutDate time.Time `json:"checkout_date" binding:"required"`
	GuestID      uint      `json:"guest_id" binding:"required"`
	RoomID       uint      `json:"room_id" binding:"required"`
	PaymentID    uint      `json:"payment_id" binding:"required"`
	Status       string    `json:"status" binding:"required"`
}
