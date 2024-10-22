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

func (r *UserRepository) Create(user models.User) error {
	query := `INSERT INTO users (name, email, password) VALUES (:name, :email, :password)`

	_, err := r.DB.NamedExec(query, user)

	if err != nil {
		return err
	}

	return nil
}
