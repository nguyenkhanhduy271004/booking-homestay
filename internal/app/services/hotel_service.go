package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	model "homestay.com/nguyenduy/internal/app/models"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/config"
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

func (s *hotelService) validateHotel(hotel *request.HotelRequest) error {
	if hotel.Name == "" {
		return errors.New("hotel name is required")
	}
	if hotel.Address == "" {
		return errors.New("hotel address is required")
	}
	if hotel.Phone == "" {
		return errors.New("hotel phone is required")
	}
	if hotel.Email == "" {
		return errors.New("hotel email is required")
	}
	if !strings.Contains(hotel.Email, "@") {
		return errors.New("invalid email format")
	}
	if hotel.Stars < 1 || hotel.Stars > 5 {
		return errors.New("stars must be between 1 and 5")
	}
	if hotel.CheckinTime.IsZero() {
		return errors.New("check-in time is required")
	}
	if hotel.CheckoutTime.IsZero() {
		return errors.New("check-out time is required")
	}
	if hotel.CheckoutTime.Before(hotel.CheckinTime) {
		return errors.New("check-out time must be after check-in time")
	}
	return nil
}

func (s *hotelService) UploadImage(file multipart.File, fileName string) (string, error) {
	tempFile, err := os.CreateTemp("", fileName)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}
	defer func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}()

	if _, err := io.Copy(tempFile, file); err != nil {
		return "", fmt.Errorf("failed to write temp file: %v", err)
	}

	cld, err := config.InitCloudinary()
	if err != nil {
		return "", fmt.Errorf("cloudinary init error: %v", err)
	}

	resp, err := cld.Upload.Upload(context.Background(), tempFile.Name(), uploader.UploadParams{
		PublicID: fileName,
		Folder:   "hotels",
	})
	if err != nil {
		return "", fmt.Errorf("cloudinary upload error: %v", err)
	}

	return resp.SecureURL, nil
}

func (s *hotelService) CreateHotel(hotel *request.HotelRequest) error {
	if err := s.validateHotel(hotel); err != nil {
		return err
	}
	return s.hotelRepo.Create(hotel)
}

func (s *hotelService) GetAllHotels() ([]model.Hotel, error) {
	return s.hotelRepo.GetAll()
}

func (s *hotelService) GetHotelByID(id uint) (*model.Hotel, error) {
	if id == 0 {
		return nil, errors.New("invalid hotel ID")
	}
	return s.hotelRepo.GetByID(id)
}

func (s *hotelService) UpdateHotel(id uint, hotel *request.HotelRequest) error {
	if id == 0 {
		return errors.New("invalid hotel ID")
	}
	if err := s.validateHotel(hotel); err != nil {
		return err
	}
	return s.hotelRepo.Update(id, hotel)
}

func (s *hotelService) DeleteHotel(id uint) error {
	if id == 0 {
		return errors.New("invalid hotel ID")
	}
	return s.hotelRepo.Delete(id)
}
