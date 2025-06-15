package services

import (
	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/request"
)

type RoomTypeService interface {
	CreateRoomType(roomType *request.RoomTypeRequest) error
	GetAllRoomTypes() ([]model.RoomType, error)
	GetRoomTypeByID(id uint) (*model.RoomType, error)
	UpdateRoomType(id uint, roomType *request.RoomTypeRequest) error
	DeleteRoomType(id uint) error
}

type roomTypeService struct {
	roomTypeRepo repository.RoomTypeRepository
}

func NewRoomTypeService(roomTypeRepo repository.RoomTypeRepository) RoomTypeService {
	return &roomTypeService{
		roomTypeRepo: roomTypeRepo,
	}
}

func (s *roomTypeService) CreateRoomType(roomType *request.RoomTypeRequest) error {
	return s.roomTypeRepo.Create(roomType)
}

func (s *roomTypeService) GetAllRoomTypes() ([]model.RoomType, error) {
	return s.roomTypeRepo.GetAll()
}

func (s *roomTypeService) GetRoomTypeByID(id uint) (*model.RoomType, error) {
	return s.roomTypeRepo.GetByID(id)
}

func (s *roomTypeService) UpdateRoomType(id uint, roomType *request.RoomTypeRequest) error {
	return s.roomTypeRepo.Update(id, roomType)
}

func (s *roomTypeService) DeleteRoomType(id uint) error {
	return s.roomTypeRepo.Delete(id)
}
