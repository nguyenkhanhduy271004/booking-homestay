package response

import "time"

type PaymentResponse struct {
	ID        uint      `json:"id"`
	BookingID uint      `json:"booking_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
