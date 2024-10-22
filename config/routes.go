package config

import (
	"github.com/darkphotonKN/gin-sqlx-template/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// base route
	api := router.Group("/api")

	// -- User --

	// --- User Setup ---
	userRepo := user.NewUserRepository(DB)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// --- User Routes ---
	userRoutes := api.Group("/users")
	userRoutes.POST("/", userHandler.CreateUserHandler)

	return router
}
