package user

import (
	"fmt"
	"net/http"

	"github.com/darkphotonKN/gin-sqlx-template/internal/models"
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
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.CreateUserService(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error when attempting to create user: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"statusCode:": http.StatusCreated, "message": "Successfully created user."})
}
