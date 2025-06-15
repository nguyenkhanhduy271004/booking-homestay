package repositories

import (
	"gorm.io/gorm"
	model "homestay.com/nguyenduy/internal/app/models"
	"homestay.com/nguyenduy/internal/request"
)

type RoomTypeRepository interface {
	Create(roomType *request.RoomTypeRequest) error
	GetAll() ([]model.RoomType, error)
	GetByID(id uint) (*model.RoomType, error)
	Update(id uint, roomType *request.RoomTypeRequest) error
	Delete(id uint) error
}

type roomTypeRepository struct {
	db *gorm.DB
}

func NewRoomTypeRepository(db *gorm.DB) RoomTypeRepository {
	return &roomTypeRepository{db: db}
}

func (r *roomTypeRepository) Create(roomType *request.RoomTypeRequest) error {
	roomTypeModel := &model.RoomType{
		Name:          roomType.Name,
		Description:   roomType.Description,
		PricePerNight: roomType.PricePerNight,
		Capacity:      roomType.Capacity,
	}
	return r.db.Create(roomTypeModel).Error
}

func (r *roomTypeRepository) GetAll() ([]model.RoomType, error) {
	var roomTypes []model.RoomType
	if err := r.db.Find(&roomTypes).Error; err != nil {
		return nil, err
	}
	return roomTypes, nil
}

func (r *roomTypeRepository) GetByID(id uint) (*model.RoomType, error) {
	var roomType model.RoomType
	if err := r.db.First(&roomType, id).Error; err != nil {
		return nil, err
	}
	return &roomType, nil
}

func (r *roomTypeRepository) Update(id uint, roomType *request.RoomTypeRequest) error {
	return r.db.Model(&model.RoomType{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":            roomType.Name,
		"description":     roomType.Description,
		"price_per_night": roomType.PricePerNight,
		"capacity":        roomType.Capacity,
	}).Error
}

func (r *roomTypeRepository) Delete(id uint) error {
	return r.db.Delete(&model.RoomType{}, id).Error
}
