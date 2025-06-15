package services

import (
	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/request"
)

type PaymentService interface {
	CreatePayment(payment *request.PaymentRequest) error
	GetAllPayments() ([]model.Payment, error)
	GetPaymentByID(id uint) (*model.Payment, error)
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

func (s *paymentService) GetAllPayments() ([]model.Payment, error) {
	return s.paymentRepo.GetAll()
}

func (s *paymentService) GetPaymentByID(id uint) (*model.Payment, error) {
	return s.paymentRepo.GetByID(id)
}

func (s *paymentService) UpdatePayment(id uint, payment *request.PaymentRequest) error {
	return s.paymentRepo.Update(id, payment)
}

func (s *paymentService) DeletePayment(id uint) error {
	return s.paymentRepo.Delete(id)
}
