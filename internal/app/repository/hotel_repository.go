package repository

import (
	"gorm.io/gorm"
	"homestay.com/nguyenduy/internal/app/model"
)

type HotelRepository interface {
	Create(hotel *model.Hotel) error
	GetAll() ([]model.Hotel, error)
	FindByID(id uint) (*model.Hotel, error)
	Update(id uint, hotel *model.Hotel) error
	Delete(id uint) error
}

type hotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) HotelRepository {
	return &hotelRepository{db: db}
}

func (r *hotelRepository) Create(hotel *model.Hotel) error {
	return r.db.Create(hotel).Error
}

func (r *hotelRepository) GetAll() ([]model.Hotel, error) {
	var hotels []model.Hotel
	err := r.db.Find(&hotels).Error
	return hotels, err
}

func (r *hotelRepository) FindByID(id uint) (*model.Hotel, error) {
	var hotel model.Hotel
	err := r.db.First(&hotel, id).Error
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (r *hotelRepository) Update(id uint, updatedHotel *model.Hotel) error {
	var hotel model.Hotel
	if err := r.db.First(&hotel, id).Error; err != nil {
		return err
	}
	return r.db.Model(&hotel).Updates(updatedHotel).Error
}

func (r *hotelRepository) Delete(id uint) error {
	return r.db.Delete(&model.Hotel{}, id).Error
}
