package converter

import (
	"homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/dtos/response"
)

func ToPaymentDTO(payment *models.Payment) response.PaymentResponse {
	return response.PaymentResponse{
		ID:        payment.ID,
		Amount:    payment.Amount,
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
	}
}
