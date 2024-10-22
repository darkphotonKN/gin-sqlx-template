package booking

import (
	"github.com/darkphotonKN/gin-sqlx-template/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BookingRepository struct {
	DB *sqlx.DB
}

func NewBookingRepository(db *sqlx.DB) *BookingRepository {
	return &BookingRepository{
		DB: db,
	}
}

func (r *BookingRepository) Create(userId uuid.UUID, req CreateBookingRequest) error {

	query := `INSERT INTO bookings(user_id, start_date, end_date, status) VALUES (:userId, :startDate, :endDate, :status)`

	// create fields required for DB insert
	fields := map[string]interface{}{
		"userId":    userId,
		"startDate": req.StartDate,
		"endDate":   req.EndDate,
		"status":    req.Status,
	}

	_, err := r.DB.NamedExec(query, fields)

	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) GetById(id uuid.UUID) (*models.Booking, error) {
	return &models.Booking{}, nil
}
