package main

import (
	"log"

	"agoravote-app-backend/src/controllers"
	"agoravote-app-backend/src/database"
	"agoravote-app-backend/src/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	database.ConnectDB()

	// Add CORS middleware
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/groups", controllers.GetGroups)
		protected.POST("/groups", controllers.CreateGroup)
		protected.GET("/user/groups", controllers.GetUserGroups)
		protected.GET("/user/profile/:id", controllers.GetUserProfile) // Include user ID as a path parameter
		protected.DELETE("/user/:id", controllers.DeleteUserAccount)   // Include user ID as a path parameter
		protected.POST("/posts", controllers.CreatePost)
		protected.POST("/votes", controllers.CreateVote)
	}

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
