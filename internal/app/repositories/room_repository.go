package repositories

import (
	"errors"

	"gorm.io/gorm"
	model "homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/dtos/request"
)

type RoomRepository interface {
	Create(room *request.RoomRequest) error
	GetAll() ([]model.Room, error)
	GetByID(id uint) (*model.Room, error)
	Update(id uint, room *request.RoomRequest) error
	Delete(id uint) error
	GetByHotelID(hotelID uint) ([]model.Room, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) Create(room *request.RoomRequest) error {
	// Check if hotel exists
	var hotel model.Hotel
	if err := r.db.First(&hotel, room.HotelID).Error; err != nil {
		return errors.New("hotel not found")
	}

	// Check if room type exists
	var roomType model.RoomType
	if err := r.db.First(&roomType, room.TypeID).Error; err != nil {
		return errors.New("room type not found")
	}

	roomModel := &model.Room{
		HotelID: hotel.ID,
		Hotel:   hotel,
		TypeID:  roomType.ID,
		Type:    roomType,
		Status:  room.Status,
	}
	return r.db.Create(roomModel).Error
}

func (r *roomRepository) GetAll() ([]model.Room, error) {
	var rooms []model.Room
	if err := r.db.Preload("Hotel").
		Preload("Type").
		Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *roomRepository) GetByID(id uint) (*model.Room, error) {
	var room model.Room
	if err := r.db.Preload("Hotel").
		Preload("Type").
		First(&room, id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) Update(id uint, room *request.RoomRequest) error {
	// Check if hotel exists
	var hotel model.Hotel
	if err := r.db.First(&hotel, room.HotelID).Error; err != nil {
		return errors.New("hotel not found")
	}

	// Check if room type exists
	var roomType model.RoomType
	if err := r.db.First(&roomType, room.TypeID).Error; err != nil {
		return errors.New("room type not found")
	}

	return r.db.Model(&model.Room{}).Where("id = ?", id).Updates(map[string]interface{}{
		"hotel_id": hotel.ID,
		"type_id":  roomType.ID,
		"status":   room.Status,
	}).Error
}

func (r *roomRepository) Delete(id uint) error {
	// Check if room has any bookings
	var count int64
	if err := r.db.Model(&model.Booking{}).Where("room_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete room with existing bookings")
	}

	return r.db.Delete(&model.Room{}, id).Error
}

func (r *roomRepository) GetByHotelID(hotelID uint) ([]model.Room, error) {
	var rooms []model.Room
	if err := r.db.Preload("Hotel").
		Preload("Type").
		Where("hotel_id = ?", hotelID).
		Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}
