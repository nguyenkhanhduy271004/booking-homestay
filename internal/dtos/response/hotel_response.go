package response

import "time"

type HotelResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Stars        int       `json:"stars"`
	Image        string    `json:"image"`
	CheckinTime  time.Time `json:"checkin_time"`
	CheckoutTime time.Time `json:"checkout_time"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
