package services

import (
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/converter"
	"homestay.com/nguyenduy/internal/dtos/request"
	"homestay.com/nguyenduy/internal/dtos/response"
)

type PaymentService interface {
	CreatePayment(payment *request.PaymentRequest) error
	GetAllPayments() ([]response.PaymentResponse, error)
	GetPaymentByID(id uint) (*response.PaymentResponse, error)
	UpdatePayment(id uint, payment *request.PaymentRequest) error
	DeletePayment(id uint) error
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepo: paymentRepo,
	}
}

func (s *paymentService) CreatePayment(payment *request.PaymentRequest) error {
	return s.paymentRepo.Create(payment)
}

func (s *paymentService) GetAllPayments() ([]response.PaymentResponse, error) {
	payments, err := s.paymentRepo.GetAll()
	if err != nil {
		return nil, err
	}

	paymentDTOs := make([]response.PaymentResponse, len(payments))
	for i, payment := range payments {
		paymentDTOs[i] = converter.ToPaymentDTO(&payment)
	}
	return paymentDTOs, nil
}

func (s *paymentService) GetPaymentByID(id uint) (*response.PaymentResponse, error) {

	payment, err := s.paymentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	paymentDTO := converter.ToPaymentDTO(payment)
	return &paymentDTO, nil
}

func (s *paymentService) UpdatePayment(id uint, payment *request.PaymentRequest) error {
	return s.paymentRepo.Update(id, payment)
}

func (s *paymentService) DeletePayment(id uint) error {
	return s.paymentRepo.Delete(id)
}
