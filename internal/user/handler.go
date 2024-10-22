package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	fmt.Println("Worked!")
}
