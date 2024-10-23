package user

import "github.com/darkphotonKN/gin-sqlx-template/internal/models"

type GetAllUsersReq struct {
	models.BaseDBDateModel
	Email    string           `db:"email" json:"email"`
	Name     string           `db:"name" json:"name"`
	Bookings []models.Booking `json:"bookings"`
}
