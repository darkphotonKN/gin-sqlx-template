package user

import (
	"github.com/darkphotonKN/gin-sqlx-template/internal/models"
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

func (r *UserRepository) Create(user models.User) (models.User, error) {
	query := `INSERT INTO users (email, password, created_at, updated_at) 
              VALUES ($1, $2, $3, $4) RETURNING id`

	var createdUser models.User
	err := r.DB.Get(&createdUser, query, user.Email, user.Password)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil

}
