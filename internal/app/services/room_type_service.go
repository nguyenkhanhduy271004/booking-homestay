package services

import (
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/converter"
	"homestay.com/nguyenduy/internal/dtos/request"
	"homestay.com/nguyenduy/internal/dtos/response"
)

type RoomTypeService interface {
	CreateRoomType(roomType *request.RoomTypeRequest) error
	GetAllRoomTypes() ([]response.RoomTypeResponse, error)
	GetRoomTypeByID(id uint) (*response.RoomTypeResponse, error)
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

func (s *roomTypeService) GetAllRoomTypes() ([]response.RoomTypeResponse, error) {
	roomTypes, err := s.roomTypeRepo.GetAll()
	if err != nil {
		return nil, err
	}

	roomTypeDTOs := make([]response.RoomTypeResponse, len(roomTypes))
	for i, roomType := range roomTypes {
		roomTypeDTOs[i] = converter.ToRoomTypeDTO(&roomType)
	}
	return roomTypeDTOs, nil
}

func (s *roomTypeService) GetRoomTypeByID(id uint) (*response.RoomTypeResponse, error) {
	roomType, err := s.roomTypeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	roomTypeDTO := converter.ToRoomTypeDTO(roomType)
	return &roomTypeDTO, nil
}

func (s *roomTypeService) UpdateRoomType(id uint, roomType *request.RoomTypeRequest) error {
	return s.roomTypeRepo.Update(id, roomType)
}

func (s *roomTypeService) DeleteRoomType(id uint) error {
	return s.roomTypeRepo.Delete(id)
}
