package services

import (
	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/request"
)

type HotelService interface {
	CreateHotel(hotel *request.HotelRequest) error
	GetAllHotels() ([]model.Hotel, error)
	GetHotelByID(id uint) (*model.Hotel, error)
	UpdateHotel(id uint, hotel *request.HotelRequest) error
	DeleteHotel(id uint) error
}

type hotelService struct {
	hotelRepo repository.HotelRepository
}

func NewHotelService(hotelRepo repository.HotelRepository) HotelService {
	return &hotelService{
		hotelRepo: hotelRepo,
	}
}

func (s *hotelService) CreateHotel(hotel *request.HotelRequest) error {
	return s.hotelRepo.Create(hotel)
}

func (s *hotelService) GetAllHotels() ([]model.Hotel, error) {
	return s.hotelRepo.GetAll()
}

func (s *hotelService) GetHotelByID(id uint) (*model.Hotel, error) {
	return s.hotelRepo.GetByID(id)
}

func (s *hotelService) UpdateHotel(id uint, hotel *request.HotelRequest) error {
	return s.hotelRepo.Update(id, hotel)
}

func (s *hotelService) DeleteHotel(id uint) error {
	return s.hotelRepo.Delete(id)
}
