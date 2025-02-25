package main

import (
	"log"
	"time"

	"agoravote-app-backend/src/controllers"
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/middleware"
	"agoravote-app-backend/src/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main
// Application entry point - configures and starts the server
func main() {
	// Initialize environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	database.ConnectDB()

	// Configure router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize services and controllers
	userService := services.NewUserService()
	groupService := services.NewGroupService()
	userController := controllers.NewUserController(userService)
	groupController := controllers.NewGroupController(groupService)

	// Public routes
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Group routes
		protected.POST("/groups", groupController.CreateGroup)
		protected.GET("/groups", groupController.GetGroups)
		protected.GET("/groups/:id", groupController.GetGroup)
		protected.GET("/user/groups", groupController.GetUserGroups)

		// User routes
		protected.GET("/user/profile/:id", userController.GetUserProfile)
		protected.DELETE("/user/:id", userController.DeleteUserAccount)
	}

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
