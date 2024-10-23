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

func (r *UserRepository) GetAll() ([]GetAllUsersReq, error) {
	query := `
	SELECT 
		users.id,
		users.name,
		users.email
	FROM users 
	LEFT JOIN bookings ON bookings.user_id = users.id
	`

	var results []struct {
		models.BaseDBDateModel
		Name      string     `db:"name"`
		Email     string     `db:"email"`
		BookingID *uuid.UUID `db:"booking_id"`
		StartDate *time.Time `db:"start_date"`
		EndDate   *time.Time `db:"end_date"`
		Status    *string    `db:"status"`
	}

	if err := r.DB.Select(&results, query); err != nil {
		return nil, err
	}

	fmt.Printf("Results: %+v\n", results)

	// map to hold each user
	resultsMap := make(map[uuid.UUID]GetAllUsersReq)

	// loop and inject all the related bookings
	for _, row := range results {
		var user, exists = resultsMap[row.ID]

		// check if user exists
		if !exists {
			// create index with user
			user = GetAllUsersReq{
				BaseDBDateModel: models.BaseDBDateModel{
					ID:        row.ID,
					CreatedAt: row.CreatedAt,
					UpdatedAt: row.UpdatedAt,
				},
				Name:     row.Name,
				Email:    row.Email,
				Bookings: []models.Booking{},
			}
		}

		// otherwise we just append to the bookings
		if row.BookingID != nil {
			user.Bookings = append(user.Bookings, models.Booking{
				BaseDBUserDateModel: models.BaseDBUserDateModel{
					ID: *row.BookingID,
				},
				Status:    *row.Status,
				StartDate: *row.StartDate,
				EndDate:   *row.EndDate,
			})
		}
	}

	// convert back to array
	users := make([]GetAllUsersReq, len(resultsMap))
	for _, user := range resultsMap {
		users = append(users, user)
	}

	fmt.Printf("users after map: %+v\n", users)

	return users, nil
}
