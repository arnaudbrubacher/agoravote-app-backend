package routes

import (
	"agoravote-app-backend/src/controllers"
	"agoravote-app-backend/src/services"

	"github.com/gin-gonic/gin"
)

// SetupRouter
// Configures all application routes and middleware
func SetupRouter() *gin.Engine {
	// Initialize router
	router := gin.Default()

	// Initialize services and controllers
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	// Configure routes
	router.GET("/user/profile/:id", userController.GetUserProfile)

	return router
}
