package repositories

import (
	"gorm.io/gorm"
	model "homestay.com/nguyenduy/internal/app/models"
)

type BookingRepository interface {
	Create(booking *model.Booking) error
	GetAll() ([]model.Booking, error)
	GetByID(id uint) (*model.Booking, error)
	Update(id uint, booking *model.Booking) error
	Delete(id uint) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(booking *model.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetAll() ([]model.Booking, error) {
	var bookings []model.Booking
	err := r.db.Preload("Guest").Preload("Room").Preload("Payment").Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetByID(id uint) (*model.Booking, error) {
	var booking model.Booking
	err := r.db.Preload("Guest").Preload("Room").Preload("Payment").First(&booking, id).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) Update(id uint, booking *model.Booking) error {
	return r.db.Model(&model.Booking{}).Where("id = ?", id).Updates(booking).Error
}

func (r *bookingRepository) Delete(id uint) error {
	return r.db.Delete(&model.Booking{}, id).Error
}
