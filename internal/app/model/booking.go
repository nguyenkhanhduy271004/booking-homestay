package model

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	GuestID      uint      `json:"guest_id"`
	Guest        Guest     `json:"guest" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoomID       uint      `json:"room_id"`
	Room         Room      `json:"room" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentID    uint      `json:"payment_id"`
	Payment      Payment   `json:"payment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CheckinDate  time.Time `json:"checkin_date"`
	CheckoutDate time.Time `json:"checkout_date"`
	TotalPrice   float64   `json:"total_price"`
	Status       string    `json:"status"`
}
