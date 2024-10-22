package user

import "github.com/darkphotonKN/gin-sqlx-template/internal/models"

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUserService(user models.User) (models.User, error) {
	return s.Repo.Create(user)
}
