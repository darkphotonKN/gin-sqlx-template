package user

import "github.com/darkphotonKN/gin-sqlx-template/internal/models"

type UserResponse struct {
	models.BaseDBDateModel
	Email    string           `db:"email" json:"email"`
	Name     string           `db:"name" json:"name"`
	Bookings []models.Booking `json:"bookings"`
}

type UserLoginResponse struct {
	RefreshToken     string `json:"refreshToken"`
	AccessToken      string `json:"accessToken"`
	AccessExpiresIn  int    `json:"accessExpiresIn"`
	RefreshExpiresIn int    `json:"refreshExpiresIn"`

	UserInfo *models.User `json:"userInfo"`
}

type UserLoginRequest struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
