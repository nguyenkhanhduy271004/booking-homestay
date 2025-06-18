package request

import "time"

type HotelRequest struct {
	Name         string    `json:"name" binding:"required"`
	Address      string    `json:"address" binding:"required"`
	Phone        string    `json:"phone" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Stars        int       `json:"stars"`
	Image        string    `json:"image"`
	CheckinTime  time.Time `json:"checkin_time" binding:"required"`
	CheckoutTime time.Time `json:"checkout_time" binding:"required"`
	Staffs       []int     `json:"staffs"`
}
