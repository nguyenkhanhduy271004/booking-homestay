package service

import (
	"homestay.com/nguyenduy/internal/app/model"
	"homestay.com/nguyenduy/internal/app/repository"
	"homestay.com/nguyenduy/internal/request"
)

type HotelService interface {
	CreateHotel(hotel *request.HotelRequest) error
	GetAllHotels() ([]model.Hotel, error)
	GetHotelById(id uint) (*model.Hotel, error)
	UpdateHotel(id uint, hotel *request.HotelRequest) error
	DeleteHotel(id uint) error
}

type hotelService struct {
	hotelRepository repository.HotelRepository
}

func NewHotelService(hotelRepository repository.HotelRepository) HotelService {
	return &hotelService{hotelRepository: hotelRepository}
}

func (h *hotelService) CreateHotel(hotel *request.HotelRequest) error {
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
	return h.hotelRepository.Create(hotelModel)
}

func (h *hotelService) GetAllHotels() ([]model.Hotel, error) {
	return h.hotelRepository.GetAll()
}

func (h *hotelService) GetHotelById(id uint) (*model.Hotel, error) {
	return h.hotelRepository.FindByID(id)
}

func (h *hotelService) UpdateHotel(id uint, hotel *request.HotelRequest) error {
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
	return h.hotelRepository.Update(id, hotelModel)
}

func (h *hotelService) DeleteHotel(id uint) error {
	return h.hotelRepository.Delete(id)
}
