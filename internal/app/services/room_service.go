package services

import (
	"errors"

	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/converter"
	"homestay.com/nguyenduy/internal/dtos/request"
	"homestay.com/nguyenduy/internal/dtos/response"
)

type RoomService interface {
	CreateRoom(room *request.RoomRequest) error
	GetAllRooms() ([]response.RoomResponse, error)
	GetRoomByID(id uint) (*response.RoomResponse, error)
	UpdateRoom(id uint, room *request.RoomRequest) error
	DeleteRoom(id uint) error
}

type roomService struct {
	roomRepo repository.RoomRepository
}

func NewRoomService(roomRepo repository.RoomRepository) RoomService {
	return &roomService{
		roomRepo: roomRepo,
	}
}

func (s *roomService) validateRoom(room *request.RoomRequest) error {
	if room.HotelID == 0 {
		return errors.New("hotel ID is required")
	}
	if room.TypeID == 0 {
		return errors.New("room type ID is required")
	}
	return nil
}

func (s *roomService) CreateRoom(room *request.RoomRequest) error {
	if err := s.validateRoom(room); err != nil {
		return err
	}
	return s.roomRepo.Create(room)
}

func (s *roomService) GetAllRooms() ([]response.RoomResponse, error) {
	rooms, err := s.roomRepo.GetAll()
	if err != nil {
		return nil, err
	}

	roomDTOs := make([]response.RoomResponse, len(rooms))
	for i, room := range rooms {
		roomDTOs[i] = converter.ToRoomDTO(&room)
	}
	return roomDTOs, nil
}

func (s *roomService) GetRoomByID(id uint) (*response.RoomResponse, error) {
	if id == 0 {
		return nil, errors.New("invalid room ID")
	}

	room, err := s.roomRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	roomDTO := converter.ToRoomDTO(room)
	return &roomDTO, nil
}

func (s *roomService) UpdateRoom(id uint, room *request.RoomRequest) error {
	if id == 0 {
		return errors.New("invalid room ID")
	}
	if err := s.validateRoom(room); err != nil {
		return err
	}
	return s.roomRepo.Update(id, room)
}

func (s *roomService) DeleteRoom(id uint) error {
	if id == 0 {
		return errors.New("invalid room ID")
	}
	return s.roomRepo.Delete(id)
}
