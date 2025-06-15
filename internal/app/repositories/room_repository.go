package repositories

import (
	"gorm.io/gorm"
	model "homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/request"
)

type RoomRepository interface {
	Create(room *request.RoomRequest) error
	GetAll() ([]model.Room, error)
	GetByID(id uint) (*model.Room, error)
	Update(id uint, room *request.RoomRequest) error
	Delete(id uint) error
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) Create(room *request.RoomRequest) error {
	roomModel := &model.Room{
		HotelID: room.HotelID,
		TypeID:  room.TypeID,
		Status:  room.Status,
	}
	return r.db.Create(roomModel).Error
}

func (r *roomRepository) GetAll() ([]model.Room, error) {
	var rooms []model.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *roomRepository) GetByID(id uint) (*model.Room, error) {
	var room model.Room
	if err := r.db.First(&room, id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) Update(id uint, room *request.RoomRequest) error {
	return r.db.Model(&model.Room{}).Where("id = ?", id).Updates(map[string]interface{}{
		"hotel_id": room.HotelID,
		"type_id":  room.TypeID,
		"status":   room.Status,
	}).Error
}

func (r *roomRepository) Delete(id uint) error {
	return r.db.Delete(&model.Room{}, id).Error
}
