package request

import "time"

type PaymentRequest struct {
	Amount        float64   `json:"amount" binding:"required"`
	PaymentDate   time.Time `json:"payment_date" binding:"required"`
	PaymentMethod string    `json:"payment_method" binding:"required"`
}
