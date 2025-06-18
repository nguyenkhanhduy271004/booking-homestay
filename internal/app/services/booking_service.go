package services

import (
	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/converter"
	"homestay.com/nguyenduy/internal/dtos/request"
	"homestay.com/nguyenduy/internal/dtos/response"
)

type BookingService interface {
	CreateBooking(booking *request.BookingRequest) error
	GetAllBookings() ([]response.BookingResponse, error)
	GetBookingByID(id uint) (*response.BookingResponse, error)
	UpdateBooking(id uint, booking *request.BookingRequest) error
	DeleteBooking(id uint) error
}

type bookingService struct {
	bookingRepo repository.BookingRepository
}

func NewBookingService(bookingRepo repository.BookingRepository) BookingService {
	return &bookingService{
		bookingRepo: bookingRepo,
	}
}

func (s *bookingService) CreateBooking(booking *request.BookingRequest) error {
	bookingModel := &model.Booking{
		GuestID:      booking.GuestID,
		RoomID:       booking.RoomID,
		PaymentID:    booking.PaymentID,
		CheckinDate:  booking.CheckinDate,
		CheckoutDate: booking.CheckoutDate,
		Status:       booking.Status,
	}
	return s.bookingRepo.Create(bookingModel)
}

func (s *bookingService) GetAllBookings() ([]response.BookingResponse, error) {
	bookings, err := s.bookingRepo.GetAll()
	if err != nil {
		return nil, err
	}

	bookingDTOs := make([]response.BookingResponse, len(bookings))
	for i, booking := range bookings {
		bookingDTOs[i] = converter.ToBookingDTO(&booking)
	}

	return bookingDTOs, nil
}

func (s *bookingService) GetBookingByID(id uint) (*response.BookingResponse, error) {
	booking, err := s.bookingRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	bookingDTO := converter.ToBookingDTO(booking)
	return &bookingDTO, nil
}

func (s *bookingService) UpdateBooking(id uint, booking *request.BookingRequest) error {
	bookingModel := &model.Booking{
		GuestID:      booking.GuestID,
		RoomID:       booking.RoomID,
		PaymentID:    booking.PaymentID,
		CheckinDate:  booking.CheckinDate,
		CheckoutDate: booking.CheckoutDate,
		Status:       booking.Status,
	}
	return s.bookingRepo.Update(id, bookingModel)
}

func (s *bookingService) DeleteBooking(id uint) error {
	return s.bookingRepo.Delete(id)
}
