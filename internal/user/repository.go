package user

import (
	"fmt"
	"time"

	"github.com/darkphotonKN/gin-sqlx-template/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user models.User) error {
	query := `INSERT INTO users (name, email, password) VALUES (:name, :email, :password)`

	_, err := r.DB.NamedExec(query, user)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetById(id uuid.UUID) (*models.User, error) {
	query := `SELECT * FROM users WHERE users.id = $1`

	var user models.User

	err := r.DB.Get(&user, query, id)

	if err != nil {
		return nil, err
	}

	// Remove password from the struct
	user.Password = ""

	return &user, nil
}

func (r *UserRepository) GetAll() ([]*UserResponse, error) {
	query := `
	SELECT 
		users.id,
		users.name,
		users.email,
		users.created_at,
		users.updated_at,
		bookings.id as booking_id,
		bookings.start_date,
		bookings.end_date,
		bookings.status,
		bookings.created_at as booking_created_at,
		bookings.updated_at as booking_updated_at
	FROM users 
	LEFT JOIN bookings ON bookings.user_id = users.id
	`

	// temporary struct to hold flat information
	var results []struct {
		models.BaseDBDateModel
		Name             string     `db:"name"`
		Email            string     `db:"email"`
		BookingID        *uuid.UUID `db:"booking_id"`
		StartDate        *time.Time `db:"start_date"`
		EndDate          *time.Time `db:"end_date"`
		Status           *string    `db:"status"`
		BookingCreatedAt *time.Time `db:"booking_created_at"`
		BookingUpdatedAt *time.Time `db:"booking_updated_at"`
	}

	if err := r.DB.Select(&results, query); err != nil {
		return nil, err
	}

	fmt.Printf("Results: %+v\n", results)

	// map to hold each user
	usersMap := make(map[uuid.UUID]*UserResponse)

	// loop and inject all the related bookings
	for _, row := range results {
		var user, exists = usersMap[row.ID]

		// check if user exists
		if !exists {
			// create index with user
			user = &UserResponse{
				BaseDBDateModel: models.BaseDBDateModel{
					ID:        row.ID,
					CreatedAt: row.CreatedAt,
					UpdatedAt: row.UpdatedAt,
				},
				Name:     row.Name,
				Email:    row.Email,
				Bookings: []models.Booking{},
			}
			usersMap[row.ID] = user
		}

		// otherwise we just append to the bookings
		if row.BookingID != nil {
			user.Bookings = append(user.Bookings, models.Booking{
				BaseDBUserDateModel: models.BaseDBUserDateModel{
					ID:        *row.BookingID,
					CreatedAt: *row.BookingCreatedAt,
					UpdatedAt: *row.BookingUpdatedAt,
				},
				UserID:    row.ID,
				Status:    *row.Status,
				StartDate: *row.StartDate,
				EndDate:   *row.EndDate,
			})
		}

		usersMap[row.ID] = user
	}

	fmt.Println("usersMap:", usersMap)

	// convert back to array
	usersResponse := make([]*UserResponse, 0, len(usersMap))
	for _, user := range usersMap {
		usersResponse = append(usersResponse, user)
	}

	fmt.Printf("usersResponse after map: %+v\n", usersResponse)

	return usersResponse, nil
}
