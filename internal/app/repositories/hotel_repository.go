package repositories

import (
	"errors"

	"gorm.io/gorm"
	model "homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/dtos/request"
)

type HotelRepository interface {
	Create(hotel *request.HotelRequest) error
	GetAll() ([]model.Hotel, error)
	GetByID(id uint) (*model.Hotel, error)
	Update(id uint, hotel *request.HotelRequest) error
	Delete(id uint) error
}

type hotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) HotelRepository {
	return &hotelRepository{db: db}
}

func (r *hotelRepository) Create(hotel *request.HotelRequest) error {
	hotelModel := &model.Hotel{
		Name:         hotel.Name,
		Address:      hotel.Address,
		Phone:        hotel.Phone,
		Email:        hotel.Email,
		Stars:        hotel.Stars,
		Image:        hotel.Image,
		CheckinTime:  hotel.CheckinTime,
		CheckoutTime: hotel.CheckoutTime,
	}
	return r.db.Create(hotelModel).Error
}

func (r *hotelRepository) GetAll() ([]model.Hotel, error) {
	var hotels []model.Hotel
	if err := r.db.Preload("Staffs").
		Preload("Staffs.User").
		Preload("Staffs.User.Role").
		Find(&hotels).Error; err != nil {
		return nil, err
	}
	return hotels, nil
}

func (r *hotelRepository) GetByID(id uint) (*model.Hotel, error) {
	var hotel model.Hotel
	if err := r.db.Preload("Staffs").
		Preload("Staffs.User").
		Preload("Staffs.User.Role").
		First(&hotel, id).Error; err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (r *hotelRepository) Update(id uint, hotel *request.HotelRequest) error {
	return r.db.Model(&model.Hotel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":          hotel.Name,
		"address":       hotel.Address,
		"phone":         hotel.Phone,
		"email":         hotel.Email,
		"stars":         hotel.Stars,
		"image":         hotel.Image,
		"checkin_time":  hotel.CheckinTime,
		"checkout_time": hotel.CheckoutTime,
	}).Error
}

func (r *hotelRepository) Delete(id uint) error {
	// Check if hotel has any staff
	var count int64
	if err := r.db.Model(&model.Staff{}).Where("hotel_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete hotel with existing staff")
	}

	// Check if hotel has any rooms
	if err := r.db.Model(&model.Room{}).Where("hotel_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete hotel with existing rooms")
	}

	return r.db.Delete(&model.Hotel{}, id).Error
}
