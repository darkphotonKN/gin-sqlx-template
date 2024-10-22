package booking

import (
	"github.com/darkphotonKN/gin-sqlx-template/internal/models"
	"github.com/google/uuid"
)

type BookingService struct {
	Repo *BookingRepository
}

func NewBookingService(repo *BookingRepository) *BookingService {
	return &BookingService{
		Repo: repo,
	}
}

func (s *BookingService) GetBookingByIdService(id uuid.UUID) (*models.Booking, error) {
	// return s.Repo.GetById(id)
	return &models.Booking{}, nil
}

func (s *BookingService) CreateBookingService(userId uuid.UUID, req CreateBookingRequest) error {
	return s.Repo.Create(userId, req)
}
