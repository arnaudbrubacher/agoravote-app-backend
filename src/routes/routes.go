package routes

import (
	"agoravote-app-backend/src/controllers"
	"agoravote-app-backend/src/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	router.GET("/user/profile/:id", userController.GetUserProfile)

	return router
}
