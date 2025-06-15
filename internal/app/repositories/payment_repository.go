package repositories

import (
	"gorm.io/gorm"
	model "homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/request"
)

type PaymentRepository interface {
	Create(payment *request.PaymentRequest) error
	GetAll() ([]model.Payment, error)
	GetByID(id uint) (*model.Payment, error)
	Update(id uint, payment *request.PaymentRequest) error
	Delete(id uint) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(payment *request.PaymentRequest) error {
	paymentModel := &model.Payment{
		Amount:        payment.Amount,
		PaymentDate:   payment.PaymentDate,
		PaymentMethod: payment.PaymentMethod,
	}
	return r.db.Create(paymentModel).Error
}

func (r *paymentRepository) GetAll() ([]model.Payment, error) {
	var payments []model.Payment
	if err := r.db.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *paymentRepository) GetByID(id uint) (*model.Payment, error) {
	var payment model.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) Update(id uint, payment *request.PaymentRequest) error {
	return r.db.Model(&model.Payment{}).Where("id = ?", id).Updates(map[string]interface{}{
		"amount":         payment.Amount,
		"payment_date":   payment.PaymentDate,
		"payment_method": payment.PaymentMethod,
	}).Error
}

func (r *paymentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Payment{}, id).Error
}
