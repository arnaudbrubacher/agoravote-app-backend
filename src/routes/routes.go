package routes

import (
	"agoravote-app-backend/src/controllers"
	"agoravote-app-backend/src/middleware"
	"agoravote-app-backend/src/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter
// Configures all application routes and middleware
func SetupRouter() *gin.Engine {
	// Initialize router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Initialize services and controllers
	userService := services.NewUserService()
	groupService := services.NewGroupService()
	userController := controllers.NewUserController(userService)
	groupController := controllers.NewGroupController(groupService)

	// Public routes
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.Signup)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		protected.GET("/user/profile/:id", userController.GetUserProfile)
		protected.DELETE("/user/:id", userController.DeleteUserAccount)

		// Group routes
		protected.POST("/groups", groupController.CreateGroup)
		protected.GET("/groups", groupController.GetGroups)
		protected.GET("/groups/:id", groupController.GetGroup)
		protected.GET("/user/groups", groupController.GetUserGroups)
		protected.POST("/groups/:id/invite", groupController.InviteToGroup)
		protected.POST("/groups/join/:token", groupController.AcceptInvitation)
		protected.PUT("/groups/:id", groupController.UpdateGroup)
		protected.POST("/groups/:id/members", groupController.AddMember)
		protected.DELETE("/groups/:id/members/:memberId", groupController.RemoveMember)
	}

	return router
}
