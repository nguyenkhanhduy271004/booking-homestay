package services

import (
	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/request"
)

type RoomService interface {
	CreateRoom(room *request.RoomRequest) error
	GetAllRooms() ([]model.Room, error)
	GetRoomByID(id uint) (*model.Room, error)
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

func (s *roomService) CreateRoom(room *request.RoomRequest) error {
	return s.roomRepo.Create(room)
}

func (s *roomService) GetAllRooms() ([]model.Room, error) {
	return s.roomRepo.GetAll()
}

func (s *roomService) GetRoomByID(id uint) (*model.Room, error) {
	return s.roomRepo.GetByID(id)
}

func (s *roomService) UpdateRoom(id uint, room *request.RoomRequest) error {
	return s.roomRepo.Update(id, room)
}

func (s *roomService) DeleteRoom(id uint) error {
	return s.roomRepo.Delete(id)
}
